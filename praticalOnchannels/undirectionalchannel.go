package main

import (
	"fmt"
	"time"
)

// Function that sends logs (send-only channel)
func logProducer(logs chan<- string) {
	for i := 1; i <= 5; i++ {
		logs <- fmt.Sprintf("Log Message %d", i)
		time.Sleep(time.Millisecond * 500) // Simulate delay
	}
	close(logs) // Close the channel when done
}

// Function that receives logs (receive-only channel)
func logConsumer(logs <-chan string) {
	for log := range logs {
		fmt.Println("Received:", log)
	}
}

func main() {
	logChannel := make(chan string) // Create a channel

	go logProducer(logChannel) // Start producer
	logConsumer(logChannel)    // Start consumer (no goroutine needed)
}
