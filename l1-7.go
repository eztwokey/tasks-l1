package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap) Len() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return len(sm.data)
}

func main() {
	// Создаем потокобезопасную map
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id)
			value := id * 10
			safeMap.Set(key, value)
			fmt.Printf("Горутина %d записала: %s -> %d\n", id, key, value)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\nВсего элементов в map: %d\n", safeMap.Len())
	fmt.Println("Содержимое map:")

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		if value, exists := safeMap.Get(key); exists {
			fmt.Printf("%s: %d\n", key, value)
		}
	}
}
