/**
 * Problem 3: find the largest prime factor
 */

package main

import "fmt"

func assign3() {

	primeFactor := 600851475143
	largestPrime := 0

	for i := 2; primeFactor > 1; i++ {
		for (primeFactor % i) == 0 {
			if largestPrime < i {
				largestPrime = i
			}
			primeFactor = primeFactor / i
		}
	}
	fmt.Println(largestPrime)
}
