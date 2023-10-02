package embeddings

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	config2 "github.com/goylold/lowcode/config"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/utils"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

type RequestParams struct {
	Input string `json:"input" from:"input" binding:"required"`
}

func init() {
	config := openai.DefaultConfig(config2.GetConConfig().OpenAIKey)
	config.BaseURL = config2.GetConConfig().OpenAIBaseUel
	client = openai.NewClientWithConfig(config)
}

func EmbeddingService(c *gin.Context) {
	claims, isExits := c.Get("claims")
	if !isExits {
		c.String(200, "无法获取当前登录用户信息，请重新登陆！")
		return
	}
	currentUser := claims.(*common.CustomClaims).Username
	request := RequestParams{}
	c.ShouldBindJSON(&request)
	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Input: request.Input,
		Model: openai.AdaEmbeddingV2,
	})
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	tokenConsumption := models.TokenConsumption{
		Amount:  int64(utils.GetTokenCount(request.Input)),
		UserID:  currentUser,
		Message: "text-embedding-ada-002",
	}
	tokenConsumption.Add()
	common.ResultSuccess(resp, c)
}

func RouteRegister(engine *gin.Engine) {
	group := engine.Group("/api/embeddings")
	{
		group.POST("/", EmbeddingService)
	}
}
