package main

import "fmt"

func main() {
	fmt.Print(toupper("Salut"))
}

func toupper(str string) string {
	word := ""
	for index, letter := range str {
		_ = index
		if letter < 'A' || letter > 'Z' {
			word = word + string(letter)
		} else {
			letter += +32
			word = word + string(letter)
		}
	}
	return word
}
