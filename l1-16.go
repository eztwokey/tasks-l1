package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2]

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// Рекурсивно сортируем части
	if 0 < right {
		quickSort(arr[:right+1])
	}
	if left < len(arr)-1 {
		quickSort(arr[left:])
	}

	return arr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(quickSort(arr))
}
