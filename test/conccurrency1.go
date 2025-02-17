package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	setup := func() {
		fmt.Println(" Initializing configuration… ")
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is calling setup.\n", id)
			once.Do(setup)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
