package explain

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	config2 "github.com/goylold/lowcode/config"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/utils"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"strings"
)

var client *openai.Client

type RequestParams struct {
	History   []models.Message
	SessionId string
	Url       string
}

func init() {
	config := openai.DefaultConfig(config2.GetConConfig().OpenAIKey)
	config.BaseURL = config2.GetConConfig().OpenAIBaseUel
	client = openai.NewClientWithConfig(config)
}

func saveAIMessage(replyMessage string, currentUser string, sessionId string, amount int64, messageIds []string) {
	replyMessageInfo := models.Message{
		Content:   replyMessage,
		SenderID:  currentUser,
		Role:      "assistant",
		SessionID: sessionId,
	}
	replyMessageInfo.Add()
	messageIds = append(messageIds, replyMessageInfo.Id)
	// 记录消息消费字数
	tokenConsumption := models.TokenConsumption{
		Amount:  int64(utils.GetTokenCount(replyMessage)) + amount,
		UserID:  currentUser,
		Message: strings.Join(append(messageIds, replyMessageInfo.Id), ","),
	}
	err := tokenConsumption.Add()
	if err != nil {
		log.Fatalln("保存用户消耗TOKEN字数出错", err.Error(), currentUser)
	}
}

func RequestService(c *gin.Context) {
	claims, isExits := c.Get("claims")
	if !isExits {
		c.String(200, "无法获取当前登录用户信息，请重新登陆！")
		return
	}
	currentUser := claims.(*common.CustomClaims).Username
	var request RequestParams
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("ShouldBindJSON error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	if len(request.History) == 0 {
		c.String(200, "请传入一条以上消息！！")
		return
	}
	var openAIHistory []openai.ChatCompletionMessage
	var messageIds []string
	var amount int64
	openAIHistory = append(openAIHistory, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Content: "请首先在输出代码所属的编程语言，输出格式为[编程语言]二十个字以内的代码功能解释，" +
			"其次请对用户输入的代码按行进行解释，格式为： [起始行号][终止行号][代码解释文本]。从上往下按功能划分代码块，用起始行号与终止行号依次解释每一行代码或代码块的作用，简单的代码不用解释，严格按照以上要求格式输出。",
	})
	openAIHistory = append(openAIHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: ` 1. package main\n 2. func main(){\n 3. a := 3 \n 4. b := a * 2 \n 5. fmt.Println(b)\n 6. }`,
	})
	openAIHistory = append(openAIHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: `[Go]定义变量a,b并输出b的值\n[1][1]定义包名。\n[2][2]定义main函数。\n[3][4]定义变量a为2，定义b=2*a。\n[5][5]用fmt包的Println方法输出变量b。`,
	})
	for _, history := range openAIHistory {
		amount += int64(utils.GetTokenCount(history.Content))
	}
	for _, message := range request.History {
		if message.Id != "" {
			messageIds = append(messageIds, message.Id)
		}
		lines := strings.Split(message.Content, "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] = fmt.Sprintf(" %d. %s", i+1, lines[i])
		}
		content := strings.Join(lines, "\n")
		amount += int64(utils.GetTokenCount(content))
		openAIHistory = append(openAIHistory, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: content,
		})
	}
	// 获取最后一条消息存储数据库
	lastMessage := request.History[len(request.History)-1]
	lastMessage.SenderID = currentUser
	err := lastMessage.Add()
	messageIds = append(messageIds, lastMessage.Id)
	if err != nil {
		// 消息发送失败
		c.String(200, "消息发送失败！"+err.Error())
		return
	}
	useModel := openai.GPT3Dot5Turbo
	if amount > 2000 {
		useModel = openai.GPT3Dot5Turbo16K
	}
	req := openai.ChatCompletionRequest{
		Model:       useModel,
		Temperature: 0,
		Messages:    openAIHistory,
		Stream:      true,
	}
	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Println("ChatCompletionStream error:", err)
		return
	}
	defer stream.Close()
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Stream(func(w io.Writer) bool {
		replyMessage := ""
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream is finished")
				saveAIMessage(replyMessage, currentUser, lastMessage.SessionID, amount, messageIds)
				return false
			}
			if err != nil {
				fmt.Println("Stream error:", err)
				saveAIMessage(replyMessage, currentUser, lastMessage.SessionID, amount, messageIds)
				return false
			}
			replyMessage += response.Choices[0].Delta.Content
			_, err = c.Writer.WriteString(response.Choices[0].Delta.Content)
			if err != nil {
				fmt.Println("The connect is closed", err.Error())
				saveAIMessage(replyMessage, currentUser, lastMessage.SessionID, amount, messageIds)
				stream.Close()
				return false
			}
			c.Writer.Flush()
		}
		return false
	})
}

func RouteRegister(engine *gin.Engine) {
	group := engine.Group("/api/explain/")
	{
		group.POST("/request", RequestService)
	}
}
