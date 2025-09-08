package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
	mu    sync.Mutex
}

// Метод с использованием atomic
func (c *Counter) IncrementAtomic() {
	atomic.AddInt64(&c.value, 1)
}

// Метод с использованием mutex
func (c *Counter) IncrementMutex() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.IncrementAtomic() // или counter.IncrementMutex()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение:", counter.GetValue())
}
