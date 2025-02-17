package main

import "fmt"

// Function to convert a slice to a channel
func sliceToChannel(s []int) chan int {
    c := make(chan int)
    go func() {
        for _, v := range s {
            c <- v
        }
        close(c)
    }()
    return c
}

// Function to square numbers
func squareNumbers(c chan int) chan int {
    out := make(chan int)
    go func() {
        for n := range c {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Function to filter even numbers
func filterEven(c chan int) chan int {
    out := make(chan int)
    go func() {
        for n := range c {
            if n%2 == 0 {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

// Function to print the final result
func printResult(c chan int) {
    for n := range c {
        fmt.Println(n)
    }
}


func main() {
	nums := []int{2, 3, 4, 7, 1}

	// stage 1: Square numbers
	datachannel := sliceToChannel(nums)
	squarechannel := squareNumbers(datachannel)

	// stage 2: Filter even numbers
	evenchannel := filterEven(squarechannel)

    // stage 3: Print the final result
    printResult(evenchannel)
}