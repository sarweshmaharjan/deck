/**
 * Problem 9:
 */

package main

import (
	"fmt"
	"math"
)

func assign9() {
	var a, b, c int
	n := 1000
	for c = int((n/3 + 1)); c < n/2; c++ {

		sqa_b := c*c - n*n + 2*n*c
		a_b := int(math.Sqrt(float64(sqa_b)))

		if a_b*a_b == sqa_b {
			b = (n - c + a_b) / 2
			a = n - b - c
			fmt.Println(a * b * c)
		}
	}
}
