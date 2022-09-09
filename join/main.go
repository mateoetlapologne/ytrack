package main

import "fmt"

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(elems, "zizi"))
}

func Join(tableau []string, sep string) string {
	word := ""
	for index, comp := range tableau {
		if index < len(tableau)-1 {
			word += comp + sep
		} else {
			word += comp
		}
	}
	return word
}
