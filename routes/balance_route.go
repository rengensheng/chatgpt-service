package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func BalanceRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/balance")
	{
		group.POST("/list", services.BalanceList)
		group.POST("/add", services.BalanceAdd)
		group.POST("/update", services.BalanceUpdate)
		group.POST("/delete/:id", services.BalanceDelete)
		group.POST("/get/:id", services.BalanceGetOne)
	}
}
