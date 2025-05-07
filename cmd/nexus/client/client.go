package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"golang.org/x/net/websocket"
	"log"
)

type Client struct {
	conn *websocket.Conn
	send chan interface{}
}

func NewClient(url string) (*Client, error) {
	conn, err := websocket.Dial(url, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		send: make(chan interface{}),
	}, nil
}

func (c *Client) Send(data interface{}) error {
	return websocket.JSON.Send(c.conn, data)
}

func (c *Client) Receive(handler func(interface{})) {
	for {
		var message interface{}
		err := websocket.JSON.Receive(c.conn, &message)
		if err != nil {
			log.Printf("Receive error: %v", err)
			return
		}
		handler(message)
	}
}

// Usage example
func main() {
	client, err := NewClient("ws://localhost:8080/ws")
	if err != nil {
		log.Fatal(err)
	}

	// Handle incoming messages
	go func() {
		client.Receive(func(msg interface{}) {
			// Handle received message
			log.Printf("Received: %v", msg)
		})
	}()

	// Send your neural data
	data := std.Data[any]{
		Context: core.Context{
			Beat: 1,
			// ... other fields
		},
		Point: 5,
	}

	err = client.Send(data)
	if err != nil {
		log.Printf("Send error: %v", err)
	}
}
