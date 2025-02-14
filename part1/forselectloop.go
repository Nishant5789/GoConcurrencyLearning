package main

import "fmt"

func main() {
	mychannel := make(chan string, 5)
	chars := []string{"H", "e", "l", "l", "o"}

	for _, char := range chars {
        mychannel <- char
    }
	close(mychannel)

    for msg := range mychannel {
        fmt.Print(msg)
    }
}