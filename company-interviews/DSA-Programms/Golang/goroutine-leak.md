

============================ with channel ===================


package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		val := <-ch

		fmt.Println("channel value", val)
	}()

	fmt.Println("Main exits, but goroutine is leaked")
}




============================== Use With context signal ============================