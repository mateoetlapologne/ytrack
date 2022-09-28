package main

import "fmt"

func main() {
	fmt.Println(alphacount("Hello 78 World!    4455 /"))
}

func alphacount(word string) int {
	counter := 0
	for index, letter := range word {
		_ = index
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			counter++
		}

	}
	return counter
}
