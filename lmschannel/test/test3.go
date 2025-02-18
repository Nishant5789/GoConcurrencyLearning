package main

import (
	"fmt"
	"sync"
)

func main() {
    ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	
	go func ()  {
		ch <- 10
		wg.Done()
	}()	
	fmt.Print(<-ch)
		
    // select {
    // case msg := <-ch:
    //     fmt.Println("Received:", msg)
    // case <-time.After(2 * time.Second):
    //     fmt.Println("Timeout! No message received")
    // }
}
