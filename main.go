package main

import (
	"github.com/goylold/lowcode/chat"
	"github.com/goylold/lowcode/embeddings"
	"github.com/goylold/lowcode/explain"
	"github.com/goylold/lowcode/image"
	"github.com/goylold/lowcode/utils"
	"github.com/goylold/lowcode/web"
	"net/http"

	"github.com/goylold/lowcode/database"

	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/config"
	"github.com/goylold/lowcode/logger"
	"github.com/goylold/lowcode/middleware"
	routes "github.com/goylold/lowcode/routes"
	"github.com/goylold/lowcode/validator"
)

func main() {
	database.LoadDatabase()
	logger.InitLogger()
	validator.Init()
	configBase := config.GetConConfig()
	router := gin.Default()
	router.Use(middleware.Auth())
	router.StaticFS("/upload", http.Dir(utils.GetFilePath(configBase.Service.UploadDir)))
	routes.RouterRegister(router)
	chat.RouteRegister(router)
	embeddings.RouteRegister(router)
	image.RouteRegister(router)
	web.RouteRegister(router)
	explain.RouteRegister(router)
	router.Run(":" + configBase.Service.Port)
}
