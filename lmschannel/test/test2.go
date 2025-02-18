package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrease counter when done
    fmt.Println("Worker", id, "started")
    time.Sleep(1 * time.Second)
    fmt.Println("Worker", id, "finished")
}

func main() {
//     var wg sync.WaitGroup
//     for i := 1; i <= 3; i++ {
//         wg.Add(1) // Increase counter
//         go worker(i, &wg)
//     }
//     wg.Wait() // Wait for all workers to complete
	ch := make(chan int)

	go func() { // Receiver goroutine
		fmt.Println("Received:", <-ch)
	}()

	ch <- 100 // Sender (won't block because receiver exists)
}
