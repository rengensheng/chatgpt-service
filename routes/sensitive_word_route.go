package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func SensitiveWordRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/sensitiveWord")
	{
		group.POST("/list", services.SensitiveWordList)
		group.POST("/add", services.SensitiveWordAdd)
		group.POST("/update", services.SensitiveWordUpdate)
		group.POST("/delete/:id", services.SensitiveWordDelete)
		group.POST("/get/:id", services.SensitiveWordGetOne)
	}
}
