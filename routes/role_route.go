package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func RoleRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/role")
	{
		group.POST("/list", services.RoleList)
		group.POST("/add", services.RoleAdd)
		group.POST("/update", services.RoleUpdate)
		group.POST("/delete/:id", services.RoleDelete)
		group.POST("/get/:id", services.RoleGetOne)
	}
}
