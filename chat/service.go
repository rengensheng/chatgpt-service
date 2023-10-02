package chat

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

type RequestParams struct {
	//History []openai.ChatCompletionMessage
	History []models.Message
	Secret  string
}

var client *openai.Client

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

func completionService() func(c *gin.Context) {
	return func(c *gin.Context) {
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
		ctx := context.Background()
		var openAIHistory []openai.ChatCompletionMessage
		var messageIds []string
		var amount int64
		for _, message := range request.History {
			if message.Id != "" {
				messageIds = append(messageIds, message.Id)
			}
			amount += int64(utils.GetTokenCount(message.Content))
			openAIHistory = append(openAIHistory, openai.ChatCompletionMessage{
				Role:    message.Role,
				Content: message.Content,
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
		req := openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.6,
			Messages:    openAIHistory,
			Stream:      true,
		}
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
}

func RouteRegister(engine *gin.Engine) {
	group := engine.Group("/api/chat")
	{
		group.POST("completion", completionService())
	}
}
