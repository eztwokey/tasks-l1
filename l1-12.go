package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество с помощью map
	set := make(map[string]bool)

	// Добавляем элементы в множество
	for _, item := range sequence {
		set[item] = true
	}

	// Получаем уникальные элементы
	uniqueItems := make([]string, 0, len(set))
	for item := range set {
		uniqueItems = append(uniqueItems, item)
	}

	fmt.Println(uniqueItems)
}
