package main

import (
	"fmt"
)

func main() {
	nextisprime(153)
}

func isprime(nb int) bool {
	res := true
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			res = false
		}
	}
	return res

}

func nextisprime(nb int) {

	if isprime(nb) == true {
		fmt.Print(nb)
	} else {
		nextisprime(nb + 1)
	}

}
