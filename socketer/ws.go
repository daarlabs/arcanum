package socketer

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gorilla/websocket"
)

type Ws interface {
	Broadcast(bytes []byte)
	Find(id ...int) ([]Client, error)
	FindOne(id int) (Client, error)
	Serve(req *http.Request, res http.ResponseWriter, id int) error
	OnRead(fn func(bytes []byte)) Ws

	MustFind(id ...int) []Client
	MustFindOne(id int) Client
	MustServe(req *http.Request, res http.ResponseWriter, id int)
}

type ws struct {
	config   Config
	hub      *hub
	upgrader *websocket.Upgrader
	onRead   func(bytes []byte)
}

func New(config ...Config) Ws {
	var cfg Config
	if len(config) == 0 {
		cfg = defaultConfig
	}
	if len(config) > 0 {
		cfg = config[0]
	}
	h := createHub()
	go h.run()
	return &ws{
		config:   cfg,
		hub:      h,
		upgrader: createUpgrader(cfg),
	}
}

func (s *ws) Broadcast(bytes []byte) {
	s.hub.broadcast <- bytes
}

func (s *ws) Find(id ...int) ([]Client, error) {
	result := make([]Client, 0)
	for c := range s.hub.clients {
		if slices.Contains(id, c.id) {
			result = append(result, c)
		}
	}
	if len(result) == 0 {
		return result, ErrorInvalidClient
	}
	return result, nil
}

func (s *ws) MustFind(id ...int) []Client {
	c, err := s.Find(id...)
	if err != nil {
		panic(err)
	}
	return c
}

func (s *ws) FindOne(id int) (Client, error) {
	for c := range s.hub.clients {
		if c.id == id {
			return c, nil
		}
	}
	return nil, ErrorInvalidClient
}

func (s *ws) MustFindOne(id int) Client {
	c, err := s.FindOne(id)
	if err != nil {
		panic(err)
	}
	return c
}

func (s *ws) OnRead(fn func(bytes []byte)) Ws {
	s.onRead = fn
	return s
}

func (s *ws) Serve(req *http.Request, res http.ResponseWriter, id int) error {
	conn, err := s.upgrader.Upgrade(res, req, nil)
	if err != nil {
		return err
	}
	defer func() {
		if e := recover(); e != nil {
			if err := conn.Close(); err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Error(res, fmt.Sprintf("%v", e), http.StatusInternalServerError)
		}
	}()
	c := createClient(s.config, s.hub, conn, id)
	s.hub.register <- c
	go func() {
		for {
			select {
			case message := <-s.hub.read:
				if s.onRead != nil {
					s.onRead(message)
				}
			}
		}

	}()
	go c.watchRead()
	go c.watchWrite()
	return nil
}

func (s *ws) MustServe(req *http.Request, res http.ResponseWriter, id int) {
	err := s.Serve(req, res, id)
	if err != nil {
		panic(err)
	}
}
