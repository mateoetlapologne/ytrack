package main

import (
	"fmt"
)

func main() {
	fmt.Print(appendrange(5, 10))
}

func appendrange(min int, max int) []int {
	a := []int{}
	if min >= max {
		return nil
	} else {
		for i := (max - min); i < max; i++ {
			a = append(a, i)
		}
		return a
	}
}
