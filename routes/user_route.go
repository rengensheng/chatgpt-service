package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func UserRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/user")
	{
		group.POST("/list", services.UserList)
		group.POST("/add", services.UserAdd)
		group.POST("/update", services.UserUpdate)
		group.POST("/delete/:id", services.UserDelete)
		group.POST("/get/:id", services.UserGetOne)
		group.POST("/accountExist", services.AccountExist)
	}
}
