// Write a function that returns the nth rune of a string. If not possible, it returns 0. without casting
package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthRune("Hello", 3))
}

func nthRune(s string, n int) rune {
	for i, r := range s {
		if i == n {
			return r
		}
	}
	return 0
}
