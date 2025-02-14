package main

import "fmt"

func main() {
	mychannel := make(chan string)
	anotherchannel := make(chan string)

	go func() {
        mychannel <- "Hello, World!"
        close(mychannel)
    }()

	go func() {
        anotherchannel <- "Goodbye, World!"
        close(anotherchannel)
    }()

	select {
		case msg := <-mychannel:
            fmt.Println(msg)
        case msg := <-anotherchannel:
            fmt.Println(msg)
    }

    // This will block the main goroutine

	msg, ok := <-mychannel
	if !ok {
        fmt.Println(msg)
    } else {
        fmt.Println("Channel closed")
    }


}