package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {
	seen := make(map[rune]bool)
	lowerStr := strings.ToLower(s)

	for _, char := range lowerStr {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	fmt.Println(hasUniqueChars("abcd"))
	fmt.Println(hasUniqueChars("abCdefAaf"))
	fmt.Println(hasUniqueChars("aabcd"))
	fmt.Println(hasUniqueChars(""))
	fmt.Println(hasUniqueChars("Go"))
}
