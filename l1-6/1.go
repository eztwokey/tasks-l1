package main

import (
	"fmt"
	"time"
)

func worker(stopFlag *bool) {
	for !*stopFlag {
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Worker stopped")
}

func main() {
	stop := false
	go worker(&stop)

	time.Sleep(2 * time.Second)
	stop = true
	time.Sleep(100 * time.Millisecond)
}
