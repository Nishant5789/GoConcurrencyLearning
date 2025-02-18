package main

// import "fmt"

func communicate(c chan int) {
    c <- 10
    // fmt.Println(<-c)
}

func main() {
    ch := make(chan int)
    go communicate(ch)
    // fmt.Println(<-ch)
}