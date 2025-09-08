package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // избегаем переполнения

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println(binarySearch(arr, 23))  // 5
	fmt.Println(binarySearch(arr, 100)) // -1
}
