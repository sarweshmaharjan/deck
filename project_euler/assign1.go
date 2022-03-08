/**
 * Multiple of 3 or 5. Find the sum of them till 1000.
 */
package main

import "fmt"

func assign1() {
	var sum int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum, " is the sum")
}
