package main

import (
	"fmt"
	"time"
)

func someFunc(str string) {
    fmt.Println(str)
}

func main() {
	go someFunc("Hello, World1!")
	go someFunc("Hello, World2!")
	go someFunc("Hello, World3!")
	go someFunc("Hello, World4!")
	go someFunc("Hello, World5!")
	go someFunc("Hello, World6!")
	go someFunc("Hello, World7!")
	time.Sleep(1 * time.Second)
}