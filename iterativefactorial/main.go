package main

import "fmt"

func main() {
	factorial(8)
}

func factorial(nb int) {
	multi := nb
	_ = multi
	resultat := 1
	_ = resultat
	for i := 1; i <= nb; i++ {
		resultat = resultat * i
	}
	fmt.Print(resultat)
}
