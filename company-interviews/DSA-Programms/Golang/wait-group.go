// 1. Print the odd no from 1 to 100 using 4 gorouines
// 2. Use a channel to collect odd numbers instead of printing directly, then sort/print in order.

package main

import (
	"fmt"
	"sort"
	"sync"
)

func main_() {
	var wg sync.WaitGroup
	ch := make(chan int)
	n := 100
	workers := 4             // no of goroutines
	chunkSize := n / workers // 25
	for w := 0; w < workers; w++ {
		wg.Add(1)
		start := w*chunkSize + 1
		end := (w + 1) * chunkSize
		go oddNumber(start, end, &wg, ch) // Go-routines
	}

	go func() { // anonymous function for wait and close channel
		wg.Wait()
		close(ch)
	}()

	var odds []int

	for val := range ch {
		odds = append(odds, val)
	}
	sort.Ints(odds)
	fmt.Println(odds)

}

func oddNumber(start, end int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i%2 != 0 {
			ch <- i // send odd number to channel
		}
	}
}
