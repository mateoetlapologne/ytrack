package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("hel", "hello"))
}

func compare(a string, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
