package main

import (
    "fmt"
    "time"
)

func worker(id int, ch chan<- string) {
    time.Sleep(time.Duration(id) * time.Second)
    ch <- fmt.Sprintf("Worker %d finished", id)
}


func main() {
    // ch := make(chan string)

    // go func() {
    //     time.Sleep(1 * time.Second)
    //     ch <- "Hello, Go!"
    // }()

    // select {
    // case msg := <-ch:
    //     fmt.Println("Received:", msg)
    // case <-time.After(2 * time.Second):
    //     fmt.Println("Timeout! No response received.")
    // }
// --------------------------------------------------------------------------------------------//
	ch1 := make(chan string)
    ch2 := make(chan string)

    go worker(2, ch1)
    go worker(1, ch2)

    for i := 0; i < 2; i++ {
        select {
        case msg := <-ch1:
            fmt.Println(msg)
        case msg := <-ch2:
            fmt.Println(msg)
        }
    }

}
