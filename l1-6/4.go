package main

import (
	"fmt"
	"runtime"
	"time"
)

func workerWithGoexit() {
	defer fmt.Println("Defer works with Goexit!")

	time.Sleep(1 * time.Second)
	fmt.Println("Calling Goexit...")
	runtime.Goexit()
	fmt.Println("This won't be printed")
}

func main() {
	go workerWithGoexit()
	time.Sleep(2 * time.Second)
}
