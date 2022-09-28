//Write a program that prints the arguments received in the command line in ASCII order.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var array []string
	for _, v := range os.Args[1:] {
		array = append(array, v)
	}
	sortwordarr(array)
}

func sortwordarr(a []string) {
	for i := 0; i+1 < len(a); i++ {
		n1 := a[i]
		n2 := a[i+1]
		if n1 < n2 {
			a[i] = n2
			a[i+1] = n1
			sortwordarr(a)
		}

	}
	for _, v := range a {
		for _, r := range v {
			z01.PrintRune(r)
		}
	}
}
