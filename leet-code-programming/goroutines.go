package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var shareMemory = []string{}
var counter int

func main_() {
	url := []string{
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.wikipedia.org/",
	}
	for _, i := range url {
		go getWebsiteStatus(i)
		wg.Add(1)
	}
	fmt.Println("pending execution ...")
	wg.Wait()
	fmt.Println("Shared Memory Items ", shareMemory)
	fmt.Println("Completed execution...")
}
func getWebsiteStatus(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error Found", err)
	} else {
		mutex.Lock()
		shareMemory = append(shareMemory, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Multiple goroutines modifying counter concurrently
	}
}

func main_() {
	go increment()
	go increment()

	time.Sleep(time.Second)
	fmt.Println("Counter:", counter) // Unpredictable output due to race condition
}
