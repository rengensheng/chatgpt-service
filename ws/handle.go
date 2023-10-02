package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hyahm/golog"
)

func SocketHandler(c *gin.Context) {
	userId := c.Param("id")
	golog.Info("创建新连接", userId)
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			log.Panicln(closeSocketErr.Error())
		}
	}()
	WsPool[userId] = ws
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			delete(WsPool, userId)
			return
		}
		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))
		err = ws.WriteJSON(struct {
			Reply string `json:"reply"`
		}{
			Reply: "Echo...",
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}
