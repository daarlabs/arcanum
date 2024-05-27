package hub

import (
	"net/http"
	
	"github.com/gorilla/websocket"
)

type Hub struct {
	broadcast  chan []byte
	read       chan []byte
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
}

const (
	hubBufferSize    = 2048
	clientBufferSize = 256
)

func New() *Hub {
	h := &Hub{
		broadcast:  make(chan []byte),
		read:       make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
	go h.listen()
	return h
}

func (h *Hub) Connect(req *http.Request, res http.ResponseWriter) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  hubBufferSize,
		WriteBufferSize: hubBufferSize,
	}
	connection, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		return err
	}
	client := &Client{
		hub:        h,
		connection: connection,
		send:       make(chan []byte, clientBufferSize),
	}
	h.clients[client] = true
	h.register <- client
	go client.read()
	go client.write()
	return nil
}

func (h *Hub) Send(value []byte) {
	if len(h.clients) == 0 {
		return
	}
	h.broadcast <- value
}

func (h *Hub) listen() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			message = trimBytesNewlines(message)
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
