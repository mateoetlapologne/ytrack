//Write a program that prints the Latin alphabet in lowercase on a single line

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
}
