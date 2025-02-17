package main

import (
	"fmt"
	"math"
	"time"
)

func generator(start, end int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			if isPrime(num) {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for _, ch := range channels {
			for num := range ch {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func takeN(in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		count := 0
		for num := range in {
			if count >= n {
				break
			}
			out <- num
			count++
		}
		close(out)
	}()
	return out
}

func main() {
	start, end := 10, 100 
	numWorkers := 3        
	limit := 10         

	numbers := generator(start, end)

	// Fan-Out - Distribute work among multiple workers
	var workers []<-chan int
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, findPrimes(numbers))
	}

	// Fan-In - Merge results from workers
	merged := fanIn(workers...)

	// Take N results
	finalResults := takeN(merged, limit)

	// Print results
	fmt.Println("Prime Numbers:")
	for prime := range finalResults {
		fmt.Println(prime)
	}

	// Adding a small delay for better visualization in some cases
	time.Sleep(1 * time.Second)
}
