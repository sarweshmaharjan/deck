/**
 * Problem 10: Summation of primes
 */

package main

import "fmt"

func assign10() {
	var sum int
	for i := 2; i <= 2000000; i++ {
		if checkPrime(i) == 1 {
			sum += i
		}
	}
	fmt.Println(sum)
}
