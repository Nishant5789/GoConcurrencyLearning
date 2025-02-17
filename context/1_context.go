package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generator := func(dataItem string, stream chan interface{}) {
		for{
			select {
            case <-ctx.Done():
                return
            case stream <- dataItem:
                time.Sleep(time.Second)
            }
		}
	}

	infiniteApples := make(chan interface{})
	go generator("Apple", infiniteApples)
    
	infiniteOranges := make(chan interface{})
	go generator("Orange", infiniteOranges)

	infinitePeaches := make(chan interface{})
	go generator("Peaches", infinitePeaches)

	go func1(ctx, &wg, 	infiniteApples)

	func2 := genericFunc
	func3 := genericFunc

	wg.Add(1)
	go func2(ctx, &wg, infiniteOranges)	

	wg.Add(1)
	go func3(ctx, &wg, infinitePeaches)	

	wg.Wait()
}

func func1(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {
	defer parentWg.Done()
    var wg sync.WaitGroup

	dowork := func(ctx context.Context){
		defer wg.Done()
		for{
			select{
				case <-ctx.Done():
                    return
                case d, ok := <-stream:
                    if !ok {
						fmt.Println("channel closed")
						return
					}
					fmt.Printf("Received: %s\n", d.(string))
                }
			}
		}

	newCtx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
        go dowork(newCtx)	
    }

	wg.Wait()
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup, stream <-chan interface{}){
	defer wg.Done()
	for{
		select{
            case <-ctx.Done():
                return
            case d, ok := <-stream:
                if!ok {
                    fmt.Println("channel1 closed")
                    return
                }
                fmt.Printf("Received: %s\n", d.(string))
        }
	}
} 