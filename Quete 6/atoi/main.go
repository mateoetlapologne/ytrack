package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	var result int
	more := 0
	more1 := 0
	for _, each := range s {
		switch each {
		case '+':
			more++
		case '-':
			more1++
		case '0':
			result = result*10 + 0
		case '1':
			result = result*10 + 1
		case '2':
			result = result*10 + 2
		case '3':
			result = result*10 + 3
		case '4':
			result = result*10 + 4
		case '5':
			result = result*10 + 5
		case '6':
			result = result*10 + 6
		case '7':
			result = result*10 + 7
		case '8':
			result = result*10 + 8
		case '9':
			result = result*10 + 9
		default:
			return 0
		}
	}
	if more == 1 {
		result = -result
	}
	if more1 == 1 {
		result = -result
	}
	if more == 0 && more1 == 0 {
		_ = result
	} else if more >= 2 || more1 >= 2 {
		result = 0
	}
	return result

}
