package socketer

import "github.com/gorilla/websocket"

func createUpgrader(config Config) *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.WriteBufferSize,
	}
}
