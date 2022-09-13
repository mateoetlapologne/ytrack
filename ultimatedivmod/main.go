package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

}

func UltimateDivMod(a *int, b *int) {
	a2 := *a
	b2 := *b
	*a = a2 / b2
	*b = b2 % a2
}
