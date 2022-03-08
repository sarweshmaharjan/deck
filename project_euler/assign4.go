/**
 * Problem 4: Largest Palindrome product
 */

package main

import "fmt"

func reverse(number int) (int, bool) {
	var remainder, reverse int
	temp := number
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10
		if number == 0 {
			break
		}
	}
	if temp == reverse {
		return temp, true
	}

	return temp, false

}
func assign4() {
	largestPalindrome := 0

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			temp, isPalindrome := reverse(i * j)
			if isPalindrome {
				if largestPalindrome < temp {
					largestPalindrome = temp
				}
			}
		}
	}
	fmt.Println(largestPalindrome)
}
