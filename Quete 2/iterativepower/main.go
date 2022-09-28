package main

import "fmt"

func main() {
	iterativepower(3, 3)
}

func iterativepower(nb int, power int) {
	resultat := nb
	_ = resultat
	power--
	for i := power; i > 0; i-- {
		resultat = nb * resultat

	}
	fmt.Print(resultat)

}
