package main

import "fmt"

func main() {
	fmt.Print(powerrecursive(4, 2))
}

func powerrecursive(nb int, power int) int {
	if power < 1 {
		return 1
	} else {
		return nb * powerrecursive(nb, power-1)
	}
}
