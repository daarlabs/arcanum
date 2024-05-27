package socketer

type hub struct {
	clients    map[*client]bool
	read       chan []byte
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func createHub() *hub {
	return &hub{
		read:       make(chan []byte),
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.write)
			}
		case message := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.write <- message:
				default:
					close(c.write)
					delete(h.clients, c)
				}
			}
		}
	}
}
