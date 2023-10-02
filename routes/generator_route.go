package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/goylold/lowcode/services"
)

func GeneratorRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/generator")
	{
		group.POST("database", services.GenerateCodeByDatabase)
	}
}
