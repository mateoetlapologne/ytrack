package main

import "fmt"

func main() {
	fmt.Print(recursivefactorial(4))
}

func recursivefactorial(nb int) int {
	if nb == 0 {
		return 1
	} else {
		return nb * recursivefactorial(nb-1)
	}
}
