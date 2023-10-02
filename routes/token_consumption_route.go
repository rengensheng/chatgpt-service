package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func TokenConsumptionRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/tokenConsumption")
	{
		group.POST("/list", services.TokenConsumptionList)
		group.POST("/add", services.TokenConsumptionAdd)
		group.POST("/update", services.TokenConsumptionUpdate)
		group.POST("/delete/:id", services.TokenConsumptionDelete)
		group.POST("/get/:id", services.TokenConsumptionGetOne)
	}
}
