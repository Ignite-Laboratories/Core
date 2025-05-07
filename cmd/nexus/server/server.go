package main

import (
	"encoding/gob"
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

// MessageHub handles bi-directional communication
type MessageHub struct {
	// Registered connections
	connections map[*websocket.Conn]bool

	// Channel for broadcasting messages
	broadcast chan interface{}

	// Register requests
	register chan *websocket.Conn

	// Unregister requests
	unregister chan *websocket.Conn
}

func NewHub() *MessageHub {
	return &MessageHub{
		connections: make(map[*websocket.Conn]bool),
		broadcast:   make(chan interface{}),
		register:    make(chan *websocket.Conn),
		unregister:  make(chan *websocket.Conn),
	}
}

func (h *MessageHub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.connections[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				conn.Close()
			}
		case message := <-h.broadcast:
			for conn := range h.connections {
				err := websocket.JSON.Send(conn, message)
				if err != nil {
					conn.Close()
					delete(h.connections, conn)
				}
			}
		}
	}
}

// Handler for WebSocket connections
func (h *MessageHub) HandleWebSocket(ws *websocket.Conn) {
	h.register <- ws
	defer func() {
		h.unregister <- ws
	}()

	// Register types that will be sent/received
	gob.Register(std.Data[any]{})
	gob.Register(core.Context{})
	gob.Register(core.Runtime{})

	for {
		var message std.Data[any]
		err := websocket.JSON.Receive(ws, &message)
		if err != nil {
			break
		}
		fmt.Println(message)
		h.broadcast <- message
	}
}

// Server setup
func main() {
	hub := NewHub()
	go hub.Run()

	http.Handle("/ws", websocket.Handler(hub.HandleWebSocket))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
