package main

import "github.com/mateoetlapologne/ytrack"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	foreach(a)
}

func foreach(a []int) int {
	for _, number := range a {
		ytrack.Printnbr(number)

	}
}
