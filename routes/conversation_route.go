package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func ConversationRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/conversation")
	{
		group.POST("/list", services.ConversationList)
		group.POST("/add", services.ConversationAdd)
		group.POST("/update", services.ConversationUpdate)
		group.POST("/delete/:id", services.ConversationDelete)
		group.POST("/get/:id", services.ConversationGetOne)
	}
}
