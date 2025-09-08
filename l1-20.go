package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	length := len(words)

	for i := 0; i < length/2; i++ {
		words[i], words[length-1-i] = words[length-1-i], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	input := "alligator bread child"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
