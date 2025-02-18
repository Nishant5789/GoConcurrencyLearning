package main

import (
	"fmt"
	"time"
)


func main() {
	ch := make(chan int)

	// close(ch)
	// ch <- 10
	go func ()  {
		time.Sleep(2 * time.Second)
		ch <- 20
	}()

	// go func(){
	// 	time.Sleep(1 * time.Second)
	// 	ch <- 20
	// }()
	// fmt.Print(<-ch)
	fmt.Print("Received:")
}