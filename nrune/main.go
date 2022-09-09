// Write a function that returns the nth rune of a string. If not possible, it returns 0. without casting
package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthRune("Hello", 1))
}

func nthRune(s string, n int) rune {
	for index, letter := range s {
		if index == n {
			return letter
		}
	}
	return 0
}
