package main

import (
	"fmt"
)

func main() {
	fmt.Println(trimatoi("hjdg6484-jy"))
}

func trimatoi(str string) int {
	result := 0
	neg := false
	neg_block := false
	for _, letter := range str {
		if letter == '-' && neg_block == false {
			neg = true
		} else if letter >= '0' && letter <= '9' {
			result = result*10 + int(letter-'0')
			neg_block = true
		}
	}
	if neg {
		result = -result
	}
	return result
}
