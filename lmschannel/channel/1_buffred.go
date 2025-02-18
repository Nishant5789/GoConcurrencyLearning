package main

import "fmt"

func main() {
	ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    close(ch)
	// ch <- 3
    
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}