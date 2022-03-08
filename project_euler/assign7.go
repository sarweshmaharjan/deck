/**
 * Problem 7: 10001st prime
 */

package main

import (
	"fmt"
)

func assign7() {
	target := 10001
	prime := []int{}
	for i := 2; len(prime) <= target; i++ {
		if checkPrime(i) == 1 {
			prime = append(prime, i)
		}
	}
	fmt.Println(prime[target-1])
}
