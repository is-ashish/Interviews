package main

import (
	"fmt"
	"time"
)

func TrafficController(dataChan chan string, stopChan chan bool) {
	for {
		// The select statement monitors channel activity
		select {
		case msg := <-dataChan:
			// Once we have a message, we use switch to evaluate its content
			switch msg {
			case "Proceed":
				fmt.Println("🚦 Green Light")
			case "Caution":
				fmt.Println("⚠️ Yellow Light")
			default:
				fmt.Println("🛑 Red Light")
			}

		case <-stopChan:
			// This case triggers when the stop signal is received
			fmt.Println("System: Shutting down...")
			return
		}
	}
}

func main_() {
	dataChan := make(chan string)
	stopChan := make(chan bool)

	// Start the controller in the background
	go TrafficController(dataChan, stopChan)

	// Simulate sending different signals
	dataChan <- "Proceed"
	dataChan <- "Caution"
	dataChan <- "Unknown"

	// Give it a tiny bit of time to process
	time.Sleep(time.Millisecond * 10)

	// Shut it down
	stopChan <- true
}
