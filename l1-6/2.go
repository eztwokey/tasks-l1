package main

import (
	"fmt"
	"time"
)

func workerWithChan(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker stopped by channel")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	go workerWithChan(stopChan)

	time.Sleep(2 * time.Second)
	close(stopChan)
	time.Sleep(100 * time.Millisecond)
}
