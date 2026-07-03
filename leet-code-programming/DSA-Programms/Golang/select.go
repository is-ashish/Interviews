package main

import (
	"fmt"
	"time"
)

func main_() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// Question -1
// Suppose you have two channels producing values at different rates. Use select to merge them into a single channel.

// Question -2
// Implement a worker goroutine that listens on a data channel and a quit channel.
// Use select to stop the worker when a quit signal is received.
