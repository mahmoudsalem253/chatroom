package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	serverAddress := "localhost:1234"

	client, err := rpc.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	// 👇 Ask for the user's name
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Connected to chat server.")
	fmt.Println("Type your messages below. Type 'exit' to quit.\n")

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		// 👇 Prefix message with user’s name
		fullMessage := fmt.Sprintf("%s: %s", name, text)

		var reply []string
		err = client.Call("ChatServer.SendMessage", fullMessage, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range reply {
			fmt.Println(msg)
		}
		fmt.Println("--------------------\n")
	}
}
