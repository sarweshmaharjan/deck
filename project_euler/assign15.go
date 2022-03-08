/**
 * Problem 15: Lattice paths
 */

package main

import "fmt"

func assign15() {
	c := 1
	for i := 1; i <= 20; i++ {
		c = c * (20 + i) / i
	}
	fmt.Println(c)
}
