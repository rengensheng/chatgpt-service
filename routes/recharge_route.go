package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func RechargeRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/recharge")
	{
		group.POST("/list", services.RechargeList)
		group.POST("/add", services.RechargeAdd)
		group.POST("/update", services.RechargeUpdate)
		group.POST("/delete/:id", services.RechargeDelete)
		group.POST("/get/:id", services.RechargeGetOne)
	}
}
