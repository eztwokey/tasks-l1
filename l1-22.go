package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(1 << 21)
	b := big.NewInt(1 << 22)

	result := new(big.Int)

	result.Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), result.String())

	result.Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a.String(), b.String(), result.String())

	result.Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), result.String())

	result.Div(b, a)
	fmt.Printf("Деление: %s / %s = %s\n", b.String(), a.String(), result.String())

	// Для очень больших чисел
	veryBig := new(big.Int)
	veryBig.Exp(big.NewInt(2), big.NewInt(1000), nil) // 2^1000
	fmt.Printf("\nОчень большое число (2^1000): %s\n", veryBig.String())
}
