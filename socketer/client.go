package socketer

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
)

type Client interface {
	Write(bytes []byte) error

	MustWrite(bytes []byte)
}

type client struct {
	config Config
	hub    *hub
	conn   *websocket.Conn
	write  chan []byte
	id     int
}

const (
	writeDuration = 10 * time.Second
	pongDuration  = 60 * time.Second
	pingPeriod    = (pongDuration * 9) / 10
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func createClient(config Config, hub *hub, conn *websocket.Conn, id int) *client {
	return &client{
		config: config,
		hub:    hub,
		conn:   conn,
		write:  make(chan []byte, config.WriteLimit),
		id:     id,
	}
}

func (c *client) Write(bytes []byte) error {
	c.write <- bytes
	return nil
}

func (c *client) MustWrite(bytes []byte) {
	c.check(c.Write(bytes))
}

func (c *client) watchRead() {
	defer func() {
		c.hub.unregister <- c
		c.close()
	}()
	c.conn.SetReadLimit(c.config.ReadLimit)
	c.check(c.conn.SetReadDeadline(time.Now().Add(pongDuration)))
	c.conn.SetPongHandler(
		func(string) error {
			c.check(c.conn.SetReadDeadline(time.Now().Add(pongDuration)))
			return nil
		},
	)
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.close()
				return
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.read <- message
	}
}

func (c *client) watchWrite() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.close()
		return
	}()
	for {
		select {
		case message, ok := <-c.write:
			c.check(c.conn.SetWriteDeadline(time.Now().Add(writeDuration)))
			if !ok {
				c.close()
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			c.check(err)
			c.checkWithLen(w.Write(message))
			n := len(c.write)
			for i := 0; i < n; i++ {
				c.checkWithLen(w.Write(newline))
				c.checkWithLen(w.Write(<-c.write))
			}
			c.check(w.Close())
		case <-ticker.C:
			c.check(c.conn.SetWriteDeadline(time.Now().Add(writeDuration)))
			c.check(c.conn.WriteMessage(websocket.PingMessage, nil))
		}
	}
}

func (c *client) check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func (c *client) checkWithLen(_ int, err error) {
	c.check(err)
}

func (c *client) close() {
	if err := c.conn.Close(); err != nil {
		panic(err)
	}
}
