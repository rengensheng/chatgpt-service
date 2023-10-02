package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func SessionRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/session")
	{
		group.POST("/list", services.SessionList)
		group.POST("/add", services.SessionAdd)
		group.POST("/update", services.SessionUpdate)
		group.POST("/delete/:id", services.SessionDelete)
		group.POST("/get/:id", services.SessionGetOne)
	}
}
