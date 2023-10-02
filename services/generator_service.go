package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/generator"
)

func GenerateCodeByDatabase(c *gin.Context) {
	results, err := generator.GenerateByDatabase()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(results, c)
}
