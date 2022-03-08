/**
 * Problem 14: Longest Collatz sequence
 */

package main

import "fmt"

func assign14() {

	initial := 0
	size := 1000000
	maxCount := 0
	count := 0
	longChain := 0

	for j := size; j > 0; j-- {
		initial = j
		for i := 0; initial > 1; i++ {
			if initial%2 == 0 {
				initial = initial / 2
			} else {
				initial = (initial * 3) + 1
			}
			count++
		}
		if maxCount < count+1 {
			maxCount = count
			longChain = j
		}
		count = 0
	}
	fmt.Println("longChain: ", longChain)
}
