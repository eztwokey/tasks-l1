package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		groupKey := int(temp/10) * 10
		groups[groupKey] = append(groups[groupKey], temp)
	}

	fmt.Println("Группировка температур:")
	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}

	fmt.Println("\nОтсортированные группы:")
	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}
}
