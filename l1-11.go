package main

import "fmt"

// Intersection универсальная функция для любых comparable типов
func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]bool)
	for _, item := range a {
		set[item] = true
	}

	var result []T
	for _, item := range b {
		if set[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {

	intsA := []int{1, 2, 3, 4, 5}
	intsB := []int{3, 4, 5, 6, 7}
	fmt.Printf("Int пересечение: %v\n", Intersection(intsA, intsB))

	stringsA := []string{"apple", "banana", "orange"}
	stringsB := []string{"banana", "kiwi", "orange", "pear"}
	fmt.Printf("String пересечение: %v\n", Intersection(stringsA, stringsB))

	floatsA := []float64{1.1, 2.2, 3.3}
	floatsB := []float64{2.2, 3.3, 4.4}
	fmt.Printf("Float пересечение: %v\n", Intersection(floatsA, floatsB))
}
