package main

import (
	"fmt"
	"sync"
)

func main() {

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(input)

		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Printf("Генерация чисел: %v\n", numbers)

		for _, x := range numbers {
			input <- x
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(output)

		for x := range input {
			result := x * 2
			output <- result
		}
	}()

	fmt.Println("Результаты (x * 2):")
	for result := range output {
		fmt.Println(result)
	}

	wg.Wait()
	fmt.Println("Завершено")
}
