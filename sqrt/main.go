package main

import "fmt"

func main() {
	sqrt(16)
}

func sqrt(nb int) {
	for i := nb; i > 0; i-- {
		if nb == (i * i) {
			fmt.Print(i)
		}
	}
}
