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
		for j := 0; j < max; j++ {
			a[j] = j + 1
		}
	}
	return a
}
