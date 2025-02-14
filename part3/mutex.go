package main

import (
	"sync"
	"fmt"
	"time"
)

var lock sync.Mutex


func processdata(data int, wg *sync.WaitGroup, output *[]int) {
	defer wg.Done()
	lock.Lock()

    time.Sleep(time.Second * 1)  
	lock.Unlock()  // Unlock after processing to prevent race conditions with other goroutines

    *output = append(*output, data*2)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	// output := []int{}
	output := make([]int, len(input))

	for _, data := range input {			
		wg.Add(1)
        go processdata(data, &wg, &output)
    }
	wg.Wait()
	fmt.Println("Processing Time:", time.Since(start))
	fmt.Println("Processed Data:", output)
}