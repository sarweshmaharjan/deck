/**
 * Problem 12: Highly divisible triangular number
 */

package main

import (
	"fmt"
	"math"
)

func assign12() {

	currentTriangle := 1
	add := 2
	for divisorCount(currentTriangle) <= 500 {
		currentTriangle += add
		add++
	}
	fmt.Println(currentTriangle)
}

func divisorCount(n int) int {
	count := 0
	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			count += 2
		}
	}
	if n%int(math.Sqrt(float64(n))) == 0 {
		count++
	}
	return count
}
