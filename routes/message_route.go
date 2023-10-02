package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func MessageRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/message")
	{
		group.POST("/list", services.MessageList)
		group.POST("/add", services.MessageAdd)
		group.POST("/update", services.MessageUpdate)
		group.POST("/delete/:id", services.MessageDelete)
		group.POST("/get/:id", services.MessageGetOne)
	}
}
