package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	BroadCast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		BroadCast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// receive client information through the channel
		case cl := <-h.Register:
			// check if roomId exists
			if _, ok := h.Rooms[cl.RoomID]; ok {
				// call the room
				r := h.Rooms[cl.RoomID]

				//check if the user isn't already in the room
				if _, ok := r.Clients[cl.ID]; !ok {
					// add client to the room
					r.Clients[cl.ID] = cl
				}
			}

		case cl := <-h.Unregister:
			// check if roomId exists
			if _, ok := h.Rooms[cl.RoomID]; ok {
				// check if the client is in the room
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					// check if there are users in the room
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						// broadcast the user has left the room
						h.BroadCast <- &Message{
							Content:  "User left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
					// remove client from the room
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					// close the message channel of the client
					close(cl.Message)
				}
			}

		case m := <-h.BroadCast:
			// check if room exists
			if _, ok := h.Rooms[m.RoomID]; ok {
				// send message to each of the client in the room
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
