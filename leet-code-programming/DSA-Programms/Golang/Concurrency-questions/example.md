Here are three interview‑level concurrency questions in Go that build on the concepts we’ve discussed (goroutines, channels, synchronization, and ordered output):
---
⚡ Question 1: Ordered Pipeline Processing
Problem:
Implement a pipeline of three stages using goroutines and channels:
Stage 1 generates numbers from 1 to N.
Stage 2 squares each number.
Stage 3 prints the results in the same order they were generated.
Ensure that results are printed strictly in sequence (1², 2², … N²).
Avoid deadlocks and race conditions.
What it tests:
Understanding of pipeline design with channels.
Maintaining order across multiple goroutines.
Proper closing of channels.



































func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	res := make(chan int, 100)
	n := 100

	//Stage 2 squares each number.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			res <- v * v
		}
		close(res)
	}()

	// Stage 3: printer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range res {
			
			fmt.Printf("%d\t", x)
		}
	}()

	// stage -:1  from  1 - n
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
	fmt.Println("Program done")
}

---
⚡ Question 2: Timeout and Cancellation
Problem:
Write a program where a goroutine simulates a long‑running task (e.g., computing factorial of a large number).
Use context.WithTimeout or select with time.After to cancel the task if it takes longer than 2 seconds.
Ensure the program exits gracefully without leaking goroutines.
What it tests:
Knowledge of Go’s context package.
Handling cancellation and timeouts in concurrent programs.
Resource cleanup and goroutine lifecycle management.

func main() {
	// var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// n := 10

	// start long running task in goroutines

	resultChan := make(chan *big.Int, 1)
	go func() {
		resultChan <- factorialBig(50)
	}()

	select {
	case <-ctx.Done():
		// timeout or cancellation trigger
		fmt.Println("Task cancelled -:", ctx.Err())
	case result := <-resultChan:
		// Task finished before timeout
		fmt.Println("Factorial computed with", len(result.String()), "digits")
		fmt.Println("Factorial computed with", len(result.String()), "digits and ", result, "No")

	}

}

// factorialBig computes factorial using math/big
func factorialBig(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}


---
⚡ Question 3: Concurrent Map Access
Problem:
You need to count word frequencies from multiple text files concurrently.
Spawn multiple goroutines to process files.
Safely update a shared map of word counts.
Print the final aggregated result.
Ensure no race conditions occur.FURL
Optimize for performance with large input.
What it tests:
Synchronization primitives (sync.Mutex, sync.Map).
Safe concurrent access to shared state.
Combining results from multiple workers.
---

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	wordCount := make(map[string]int)
	files := []string{"location.txt", "user.txt", "vehicle.txt"}
	for _, file := range files {
		wg.Add(1)
		go processFiles(file, wordCount, &wg, &mx)
	}
	wg.Wait()
	// convert map to slice

	type wordFreq struct {
		word  string
		count int
	}

	var freqs []wordFreq

	for k, v := range wordCount {
		freqs = append(freqs, wordFreq{k, v})
	}
	// Sort by frequency (descending)
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	// fmt.Println(freqs)
	for _, v := range freqs {
		fmt.Printf("word -: %s ------- count -: %d \n", v.word, v.count)
	}
}

func processFiles(txt string, wordCount map[string]int, wg *sync.WaitGroup, mx *sync.Mutex) {

	defer wg.Done()

	file, err := os.Open(txt)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Change split function to scan words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		mx.Lock()
		wordCount[word]++
		mx.Unlock()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}



===================Efficient go use ===============

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
)

func main() {
	var mx sync.Mutex
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	wordCount := make(map[string]int)
	files := []string{"location.txt", "user.txt", "vehicle.txt"}
	go processFiles(ch1, files[0])
	go processFiles(ch2, files[1])
	go processFiles(ch3, files[2])


	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				mx.Lock()
				wordCount[v]++
				mx.Unlock()
			}
		case v1, ok := <-ch2:

			if !ok {
				ch2 = nil
			} else {
				mx.Lock()
				wordCount[v1]++
				mx.Unlock()
			}
		case v2, ok := <-ch3:
			if !ok {
				ch3 = nil
			} else {
				mx.Lock()
				wordCount[v2]++
				mx.Unlock()
			}
		}

		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}

	// convert map to slice
	type wordFreq struct {
		word  string
		count int
	}

	var freqs []wordFreq

	for k, v := range wordCount {
		freqs = append(freqs, wordFreq{k, v})
	}
	// Sort by frequency (descending)
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	// fmt.Println(freqs)
	for _, v := range freqs {
		fmt.Printf("word -: %s ------- count -: %d \n", v.word, v.count)
	}
}

func processFiles(cha chan<- string, txt string) {

	file, err := os.Open(txt)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		cha <- word
	}
	close(cha)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
