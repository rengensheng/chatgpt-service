package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	config2 "github.com/goylold/lowcode/config"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/utils"
	"github.com/sashabaranov/go-openai"
	"image/png"
	"os"
	"path"
)

var client *openai.Client

type RequestParams struct {
	Input string `json:"input" from:"input" binding:"required"`
	Size  string `json:"size" from:"size" binding:"required"`
}

var amountMap map[string]int64

func init() {
	config := openai.DefaultConfig(config2.GetConConfig().OpenAIKey)
	config.BaseURL = config2.GetConConfig().OpenAIBaseUel
	client = openai.NewClientWithConfig(config)
	amountMap = make(map[string]int64)
	amountMap[openai.CreateImageSize256x256] = 8000
	amountMap[openai.CreateImageSize512x512] = 12000
	amountMap[openai.CreateImageSize1024x1024] = 16000
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
	if request.Size != openai.CreateImageSize256x256 &&
		request.Size != openai.CreateImageSize512x512 &&
		request.Size != openai.CreateImageSize1024x1024 {
		request.Size = openai.CreateImageSize256x256
	}
	amount := amountMap[request.Size]
	resp, err := client.CreateImage(context.Background(), openai.ImageRequest{
		Prompt:         request.Input,
		Size:           request.Size,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	})
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	imgBytes, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return
	}
	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		common.ResultError(500, "PNG decode error: "+err.Error(), c)
		return
	}
	// 保存文件到上传目录
	saveFileName := utils.GetUUID() + ".png"
	file, err := os.Create(path.Join(utils.GetFilePath(config2.GetConConfig().Service.UploadDir), saveFileName))
	if err != nil {
		common.ResultError(500, "PNG save error: "+err.Error(), c)
		return
	}
	defer file.Close()
	if err := png.Encode(file, imgData); err != nil {
		common.ResultError(500, "PNG encode error: : "+err.Error(), c)
		return
	}
	tokenConsumption := models.TokenConsumption{
		Amount:  amount,
		UserID:  currentUser,
		Message: "DALL-E-2",
	}
	tokenConsumption.Add()
	common.ResultSuccess(gin.H{
		"filename": saveFileName,
	}, c)
}

func RouteRegister(engine *gin.Engine) {
	group := engine.Group("/api/image")
	{
		group.POST("/", EmbeddingService)
	}
}
