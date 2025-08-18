package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	square := make(chan int, len(arr))
	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(square)
	}()

	for res := range square {
		fmt.Println(res)
	}
}
