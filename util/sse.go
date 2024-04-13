package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Clients is a list of channels to send events to connected clients
var clients = make(map[string][]chan string)

// broadcast sends an event to all connected clients
func SSEBroadcast(roomID string, data string) {
	for _, client := range clients[roomID] {
		client <- data
	}
}

func SSEConnect(c *gin.Context, roomID string) {
	// Set the response header to indicate SSE content type
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	// Create a channel to send events to the client
	// println("Client connected")
	eventChan := make(chan string)
	if clients[roomID] == nil {
		clients[roomID] = []chan string{}
	}
	clients[roomID] = append(clients[roomID], eventChan)
	defer func() {
		for _, v := range clients[roomID] {
			if v != eventChan {
				clients[roomID] = append(clients[roomID], v)
			}
		}
		close(eventChan)
	}()
	// Listen for client close and remove the client from the list
	notify := c.Writer.CloseNotify()
	go func() {
		<-notify
		// fmt.Println("Client disconnected")
	}()
	// Continuously send data to the client
	for {
		data := <-eventChan
		// println("Sending data to client", data)
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		c.Writer.Flush()
	}
}
