package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Функция для установки бита
	setBit := func(num int64, i uint, bit int) int64 {
		if bit == 1 {
			return num | (1 << i)
		}
		return num &^ (1 << i)
	}

	num := int64(5)
	i := uint(1)

	fmt.Printf("Исходное: %d (%04s)\n", num, strconv.FormatInt(num, 2))

	result := setBit(num, i, 0)
	fmt.Printf("Бит %d -> 0: %d (%04s)\n", i, result, strconv.FormatInt(result, 2))

	result = setBit(result, i, 1)
	fmt.Printf("Бит %d -> 1: %d (%04s)\n", i, result, strconv.FormatInt(result, 2))
}
