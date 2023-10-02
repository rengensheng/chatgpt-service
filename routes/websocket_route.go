package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/ws"
)

func WebsocketRouterRegistry(engine *gin.Engine) {
	group := engine.Group("/api/ws/websocket")
	{
		group.GET("/:id", ws.SocketHandler)
	}
}
