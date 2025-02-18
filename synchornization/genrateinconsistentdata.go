package main

import (
	"sync"
	"fmt"
	// "time"
)


func processdata(data int, wg *sync.WaitGroup, output *[]int) {
	defer wg.Done()

    // Simulate some processing time
    // time.Sleep(time.Second * 1)  

    *output = append(*output, data*2)
}

func main() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	output := []int{}

	for _, data := range input {			
		wg.Add(1)
        go processdata(data, &wg, &output)
    }
	wg.Wait()
	fmt.Println("Processed Data:", output)
}