package main

import (
	"fmt"
	"sync"
	
)

var mu1, mu2 sync.Mutex

func main_() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: locks mu1 then mu2
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu2.Unlock()
		mu1.Unlock()
	}()

	// Goroutine 2: locks mu2 then mu1
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu2")
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		mu1.Unlock()
		mu2.Unlock()
	}()

	wg.Wait()
}

===
“Deadlocks occur when goroutines acquire locks in inconsistent order or forget to release them. For example, if one goroutine locks mu1 then mu2, while another locks mu2 then mu1, they’ll block each other forever. The solution is consistent lock ordering, careful use of defer Unlock(), or using channels instead of mutexes.”
===
with 
