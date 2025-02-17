package main

import (
	"fmt"
	"time"
)

// Function to prepare food in the kitchen
func kitchen(order chan string) {
	fmt.Println("Kitchen is preparing the order...")
	time.Sleep(2 * time.Second) // Simulating food preparation time
	order <- "Pizza is ready!"  // Sending the order to the waiter
}

// Function to serve food
func waiter(order chan string) {
	food := <-order // Receiving the order from the kitchen
	fmt.Println("Waiter is serving:", food)
}

func main() {
	// Creating an unbuffered channel
	orderChannel := make(chan string)

	// Kitchen prepares food
	go kitchen(orderChannel)

	// Waiter waits for food
	go waiter(orderChannel)

	// Sleep to allow goroutines to execute
	time.Sleep(3 * time.Second)
}
