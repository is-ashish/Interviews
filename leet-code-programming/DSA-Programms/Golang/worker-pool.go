// ```
// Absolutely, Ashish — this is a powerful concurrency pattern in Go. Let’s break it down step by step.

// 🧠 What Is a Worker Pool?
// A worker pool is a concurrency pattern where:
// You have a fixed number of worker goroutines.
// They pull tasks from a shared channel (task queue).
// Each worker processes tasks independently.
// The main goroutine feeds tasks into the channel.

// ===========================================
// Instead, workers dynamically pick up tasks as they become available.
// ===========================================

// 🔧 Why Use a Worker Pool?
// ✅ Efficient resource usage: You control the number of goroutines.
// ✅ Dynamic task assignment: Workers grab tasks as they’re ready.
// ✅ Scalable: Works well with large or uneven workloads.
// ✅ Avoids goroutine explosion: Prevents launching thousands of goroutines.

// 🧩 Worker Pool Example (Odd Numbers)
// Go

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// what is worker pool -
//

func main_() {

	wg := sync.WaitGroup{}
	n := 250
	workers := 3
	items := make(chan int, n)
	results := make(chan int, n)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go execution(items, results, &wg)
	}
	for j := 0; j <= 250; j++ {
		items <- j
	}
	close(items)

	go func() {
		wg.Wait()
		close(results)
	}()
	var odds []int
	for result := range results {
		odds = append(odds, result)
	}
	fmt.Println(odds)
}

func execution(items <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range items {
		if item%2 != 0 {
			results <- item
		}
	}
}

// 🧠 Real-World Analogy --------
// Imagine a pizza shop:
// Orders come in and go into a queue (tasks).
// You have 4 chefs (workers) who pick up orders as they’re ready.
// Each chef finishes their pizza and sends it to the delivery counter (results).
// Once all orders are done, the counter closes.

// Would you like me to visualize this as a diagram showing the task queue, workers, and result flow? It’s a great way to internalize the pattern.

func main_() {
	var wg sync.WaitGroup
	// ch := make(chan int)
	res := make(chan int)
	wg.Add(2) // 2 goroutines

	go func() { // worker goroutines
		worker(res)
		wg.Done()
	}()

	go func() { // Printer goroutines
		for k := range res {
			fmt.Println(k)
		}
		wg.Done()
	}()
	wg.Wait()
}

func worker(res chan int) {
	for i := 1; i <= 5; i++ {
		num := rand.Intn(100)
		res <- num
	}
	close(res)
}

// ==============================================

package main

import (
	"fmt"
	"sync"
)

type Result struct {
	id  int
	val int
}

func main() {
	var wg sync.WaitGroup
	inChan := make(chan int)
	resChan := make(chan Result) // buffered

	task := 10

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go computeFactorial(inChan, resChan, &wg)
	}

	go func() {
		for i := 1; i <= task; i++ {
			inChan <- i
		}
		close(inChan)
	}()

	go func() {
		wg.Wait()
		close(resChan)
	}()

	// collect result into slice
	results := make([]int, task)
	for r := range resChan {
		results[r.id-1] = r.val
	}

	for i, v := range results {
		fmt.Printf("%d! = %d\n", i+1, v)
	}

}

func computeFactorial(inChan chan int, resChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range inChan {
		resChan <- Result{id: v, val: getFact(v)}
	}
}

func getFact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * getFact(n-1)
}
