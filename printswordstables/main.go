package main

import (
	"github.com/01-edu/z01"
)

func main() {
	a := SplitWhiteSpaces(("Hello how are you?"))
	PrintWordStables(a)
}

func SplitWhiteSpaces(s string) []string {
	var res []string
	var timeres string
	for _, lettre := range s {
		if rune(lettre) == 32 || rune(lettre) == 9 || rune(lettre) == 10 || rune(lettre) == 27 {
			res = append(res, timeres)
			timeres = ""
		} else {
			timeres += string(lettre)
		}
	}
	res = append(res, timeres)
	return res

}

func PrintWordStables(carra []string) {
	for _, newres := range carra {
		for _, char := range string(newres) {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
