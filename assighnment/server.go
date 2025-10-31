package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

// SendMessage adds a message to the chat history and returns the updated list
func (c *ChatServer) SendMessage(msg string, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.messages = append(c.messages, msg)

	// 👇 Print message to server console
	fmt.Println("New message received:", msg)

	*reply = append([]string{}, c.messages...) // return a copy of chat history
	return nil
}

// GetMessages returns all messages
func (c *ChatServer) GetMessages(dummy string, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	*reply = append([]string{}, c.messages...)
	return nil
}

func main() {
	server := new(ChatServer)
	err := rpc.Register(server)
	if err != nil {
		fmt.Println("Error registering RPC server:", err)
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}

	fmt.Println("Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

