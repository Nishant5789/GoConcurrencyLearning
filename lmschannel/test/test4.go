package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    var wg sync.WaitGroup

    // Create 3 worker goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    // Send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // Close jobs channel so workers know no more jobs are coming

    wg.Wait() // Wait for workers to finish
    close(results)

    for res := range results {
        fmt.Println("Result:", res)
    }
}
