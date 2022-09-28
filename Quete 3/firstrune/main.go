package main

import "fmt"

func main() {
	fmt.Println(FirstRune("Hello"))
}

func FirstRune(s string) rune {
	for index, lettre := range s {
		if index == 0 {
			return lettre
		}
	}
	return 0
}
