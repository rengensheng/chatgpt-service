package ws

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/hyahm/golog"
)

var WsPool map[string]*websocket.Conn

func init() {
	WsPool = make(map[string]*websocket.Conn)
	go WsHeartbeat()
}

func WsHeartbeat() {
	golog.Info("开始心跳检测...")
	for {
		time.Sleep(time.Second * 10)
		for _, ws := range WsPool {
			err := ws.WriteJSON(struct {
				Reply string `json:"reply"`
			}{
				Reply: "Echo...",
			})
			if err != nil {
				golog.Info(err.Error())
			}
		}
	}
}

func GetUserSocketById(userId string) *websocket.Conn {
	return WsPool[userId]
}
