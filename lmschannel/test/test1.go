package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
        }
        close(ch) // Closing the channel after sending data
    }()

    for val := range ch { // Loop exits when channel is closed
        fmt.Println("Received:", val)
    }
}
