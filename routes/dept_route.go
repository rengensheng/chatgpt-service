package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func DeptRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/dept")
	{
		group.POST("/list", services.DeptList)
		group.POST("/add", services.DeptAdd)
		group.POST("/update", services.DeptUpdate)
		group.POST("/delete/:id", services.DeptDelete)
		group.POST("/get/:id", services.DeptGetOne)
	}
}
