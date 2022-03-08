package main

import (
	"fmt"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func assign5() {
	result := 1
	for j := 2; j <= 20; j++ {
		result = LCM(result, j)
	}
	fmt.Println(result)
}
