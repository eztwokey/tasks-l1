package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int, chan string, chan bool:
		fmt.Println("channel")
	default:
		fmt.Printf("unknown type: %v\n", reflect.TypeOf(v))
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
}
