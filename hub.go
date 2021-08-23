package mqttclient

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// Registered clients.
	Clients map[*Client]bool

	// Broadcast commands from the clients.
	Broadcast chan []byte
}

// NewHub create new hub
func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
	}
}

// Run run hub
func (h *Hub) Run() {
	go h.run()
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Publish)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Publish <- message:
				default:
					close(client.Publish)
					delete(h.Clients, client)
				}
			}
		}
	}
}
