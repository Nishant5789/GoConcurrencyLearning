package main

import (
	"fmt"
	"time"
)

// chat function where two users exchange messages
func chat(user1 chan string, user2 chan string) {
	go func() {
		user1 <- "Hello from User 1!" // Send message from User 1
		fmt.Println("User 1 sent a message")
	}()

	time.Sleep(1 * time.Second) // Simulate delay

	go func() {
		msg := <-user1 // Receive message at User 2
		fmt.Println("User 2 received:", msg)

		user2 <- "Hello from User 2!" // Send response back
		fmt.Println("User 2 sent a message")
	}()

	time.Sleep(1 * time.Second) // Simulate delay

	go func() {
		msg := <-user2 // Receive response at User 1
		fmt.Println("User 1 received:", msg)
	}()
}

func main() {
	// Creating bidirectional channels
	user1Channel := make(chan string)
	user2Channel := make(chan string)

	chat(user1Channel, user2Channel)

	// Sleep to allow goroutines to execute
	time.Sleep(3 * time.Second)
}
