package main

import (
	"fmt"
	"sync"
)

func main__() {
	pipe := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			pipe <- i
			fmt.Printf("Item produce -: %d\n", i)
		}
		close(pipe)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for p := range pipe {
			fmt.Printf("Item consume -: %d\n", p)

		}
	}()
	wg.Wait()

}
