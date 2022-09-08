package main

import "fmt"

func main() { fmt.Print(fibonacci(3)) }

func fibonacci(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return fibonacci(index-1) + fibonacci(index-2)
	}

}
