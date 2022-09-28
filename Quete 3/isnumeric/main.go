package main

import "fmt"

func main() {
	fmt.Print(isnumeric(" 6"))
}

func isnumeric(letter string) bool {
	res := true
	for index, nbr := range letter {
		_ = index
		if nbr < '0' || nbr > '9' {
			res = false
		}
	}
	return res
}
