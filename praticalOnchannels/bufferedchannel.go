package main

import (
	"fmt"
	"time"
)

func worker(id int, taskQueue chan string) {
	for task := range taskQueue { // Continuously listen for tasks
		fmt.Printf("Worker %d processing task: %s\n", id, task)
		time.Sleep(1 * time.Second) // Simulate processing time
	}
}

func main() {
	taskQueue := make(chan string, 5) // Buffered channel with capacity 5

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, taskQueue)
	}

	// Producer: Adding 10 tasks
	tasks := []string{"Task1", "Task2", "Task3", "Task4", "Task5", "Task6", "Task7", "Task8", "Task9", "Task10"}
	for _, task := range tasks {
		fmt.Println("Adding:", task)
		taskQueue <- task // Sending tasks to the buffered channel
	}

	close(taskQueue) // Close the channel after sending all tasks
	time.Sleep(5 * time.Second) // Wait for workers to finish processing
}
