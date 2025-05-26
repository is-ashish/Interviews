package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan int)
	go func() { ch <- 42 }()
	// ch <- 42 // Blocks forever! No receiver is ready
	fmt.Println(<-ch)
}

//Channels help avoid race conditions, but improper handling can cause deadlocks:
// 8. Deadlocks & Avoiding Race Conditions

// Worker pools optimize parallelism using channels.
// Worker Pools (Efficient Task Processing)
func main_6() {
	const numWorkers = 3
	jobs := make(chan int, 10)

	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobs {
				fmt.Println("Processing job:", job)
			}
		}()
	}

	for j := 1; j <= 10; j++ {
		jobs <- j // Sending tasks
		fmt.Println(<-jobs)
	}
	close(jobs)

}

// Multiple channels feeding into one common output channel.
func main_5() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { ch1 <- 1 }()
	go func() { ch2 <- 2 }()

	output := make(chan int)

	go func() { output <- <-ch1 }()
	go func() { output <- <-ch2 }()

	fmt.Println(<-output) // Receives data from ch1 or ch2
	fmt.Println(<-output) // Receives data from ch1 or ch2

}

// Multiple goroutines reading from one channel to process tasks.
func main_2() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(<-ch) // Multiple goroutines consuming data
		}()
	}

}

// 🔹 select is used for handling multiple channels simultaneously.
func main_1() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() { ch1 <- 42 }()
	go func() { ch2 <- "Hello" }()

	select {
	case num := <-ch1:
		fmt.Println("Received from ch1:", num)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	}

}

func main_() {
	defer func() {
		fmt.Println("world")
	}()
	fmt.Println("hello")
	log.Fatal("oops")
}

// func process(i int, wg *sync.WaitGroup) {
// 	fmt.Println("started Goroutine ", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Goroutine %d ended\n", i)
// 	wg.Done()
// }
// func main_() {
// 	no := 3
// 	var wg sync.WaitGroup
// 	for i := 0; i < no; i++ {
// 		wg.Add(1)
// 		go process(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("All go routines finished executing")
// }

// func main_() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	ch := make(chan int)
// 	go func() {
// 		for {
// 			foo, ok := <-ch
// 			if !ok {
// 				println("Goroutine Closed")
// 				wg.Done()
// 				return
// 			}
// 			println(foo)
// 		}
// 	}()
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	close(ch)

// 	wg.Wait()
// }

// array sum calculator

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum = sum + v
// 	}
// 	c <- sum
// }

// func main_() {
// 	s := []int{7, 2, 8, -9, 4, 0}
// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c
// 	fmt.Println(x, y, x+y)
// }
