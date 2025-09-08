package main

import "fmt"

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	// Сдвигаем хвост слайса на место удаляемого элемента
	copy(slice[i:], slice[i+1:])

	// Обнуляем последний элемент (для предотвращения утечки памяти)
	slice[len(slice)-1] = 0

	// Уменьшаем длину слайса на 1
	return slice[:len(slice)-1]
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный:", slice)

	slice = removeElement(slice, 2)
	fmt.Println("После удаления:", slice)

	slice = removeElement(slice, 10)
	fmt.Println("Некорректный индекс:", slice)
}
