package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func LoginRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/login")
	{
		group.POST("/", services.UserLogin)
	}
	engine.POST("/api/getUserInfo", services.GetUserInfo)
	engine.POST("/api/getPermCode", services.GetPermCode)
	engine.POST("/api/getMenuList", services.GetMenuList)
}
