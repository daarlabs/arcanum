package hub

import (
	"time"
	
	"github.com/gorilla/websocket"
)

type Client struct {
	hub        *Hub
	connection *websocket.Conn
	send       chan []byte
}

const (
	pingPeriod = time.Second * 60
)

func (c *Client) Send(value []byte) {
	c.send <- value
}

func (c *Client) read() {
	defer func() {
		if ok := c.hub.clients[c]; ok {
			c.hub.unregister <- c
			if err := c.connection.Close(); err != nil {
				return
			}
		}
	}()
	for {
		_, message, err := c.connection.ReadMessage()
		if err != nil && !websocket.IsUnexpectedCloseError(err) {
			return
		}
		message = trimBytesNewlines(message)
		c.hub.read <- message
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		if ok := c.hub.clients[c]; ok {
			if err := c.connection.Close(); err != nil {
				return
			}
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
