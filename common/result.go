package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResultError(code int, message interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"type":    "fail",
	})
}

func ResultSuccess(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"result":  data,
		"type":    "success",
	})
}

func ResultSuccessList(data interface{}, total int64, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"result": gin.H{
			"items": data,
			"total": total,
		},
		"type": "success",
	})
}
