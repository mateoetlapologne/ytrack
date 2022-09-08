//Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third. These combinations are separated by a comma and a space

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	PrintComb()
}

func PrintComb() {

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				z01.PrintRune(rune(i + 48))
				z01.PrintRune(rune(j + 48))
				z01.PrintRune(rune(k + 48))

				if i != 7 {
					z01.PrintRune(',')
					z01.PrintRune(32)
				}
			}
		}
	}
	fmt.Println()
}
