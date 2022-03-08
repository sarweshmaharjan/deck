/**
 * Problem 2: Even fibonacci number.
 */
package main

import "fmt"

func assign2() {

	fib := []int{1, 2}
	var sum int

	for _, no := range fib {
		if no%2 == 0 {
			sum = sum + no
		}
	}
	for i := 2; fib[i-1] <= 4000000; i++ {
		if (fib[i-1]+fib[i-2])%2 == 0 {
			sum = sum + fib[i-1] + fib[i-2]
		}
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println(sum)
}
