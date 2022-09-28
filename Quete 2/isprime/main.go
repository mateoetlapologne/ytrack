package main

import "fmt"

func main() {
	isprime(152)
}

func isprime(nb int) bool {
	res := true
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			res = false
		}
	}
	fmt.Print(res)
	return res
}
