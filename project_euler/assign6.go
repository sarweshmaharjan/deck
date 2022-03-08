/**
 * Problem 6: Sum square difference
 */

package main

import "fmt"

func assign6() {
	var sumOfSquare, squareOfSum int
	for i := 1; i <= 100; i++ {
		sumOfSquare = sumOfSquare + (i * i)
		squareOfSum = squareOfSum + i
	}
	squareOfSum = squareOfSum * squareOfSum
	difference := squareOfSum - sumOfSquare
	fmt.Println(difference)
}
