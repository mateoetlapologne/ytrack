/* Write a program that is called doop.

The program has to be used with three arguments:

A value
An operator, one of : +, -, /, *, %
Another value
In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below. */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[3])
		if os.Args[2] == "+" {
			fmt.Println(a + b)
		} else if os.Args[2] == "-" {
			fmt.Println(a - b)
		} else if os.Args[2] == "*" {
			fmt.Println(a * b)
		} else if os.Args[2] == "/" {
			if b != 0 {
				fmt.Println(a / b)
			}
		} else if os.Args[2] == "%" {
			if b != 0 {
				fmt.Println(a % b)
			}
		}
	}
}
