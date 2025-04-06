package main

import (
	"fmt"
	"log"
)

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
