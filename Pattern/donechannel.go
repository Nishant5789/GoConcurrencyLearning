package main

import (
	"fmt"
	"time"
)

func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Print("...")
		}
	}
}

func main() {
	done := make(chan bool)
	go dowork(done)
	time.Sleep(2 * time.Second)
	close(done)
}