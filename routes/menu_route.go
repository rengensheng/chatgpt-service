package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func MenuRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/menu")
	{
		group.POST("/list", services.MenuList)
		group.POST("/add", services.MenuAdd)
		group.POST("/update", services.MenuUpdate)
		group.POST("/delete/:id", services.MenuDelete)
		group.POST("/get/:id", services.MenuGetOne)
	}
}
