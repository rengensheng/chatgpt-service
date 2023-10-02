package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func WebUserRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/webUser")
	{
		group.POST("/list", services.WebUserList)
		group.POST("/add", services.WebUserAdd)
		group.POST("/update", services.WebUserUpdate)
		group.POST("/delete/:id", services.WebUserDelete)
		group.POST("/get/:id", services.WebUserGetOne)
		group.POST("/login", services.WebUserLogin)
		group.POST("/register", services.WebUserRegister)
	}
}
