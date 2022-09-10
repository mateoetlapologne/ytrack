package main

import (
	"fmt"
)

func main() {
	fmt.Print(makerange(5, 10))
}

func makerange(min int, max int) []int {
	a := make([]int, (max - min))
	if min >= max {
		return nil
	} else {
		for j := 0; j < max-min; j++ {
			a[j] = min + j
		}
	}
	return a
}
