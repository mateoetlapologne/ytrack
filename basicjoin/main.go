package main

import "fmt"

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}

func BasicJoin(tableau []string) string {
	word := ""
	for _, comp := range tableau {
		word += comp
	}
	return word
}
