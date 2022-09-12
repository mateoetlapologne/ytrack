package main

import "fmt"

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsNumeric(letter string) bool {
	res := true
	for index, nbr := range letter {
		_ = index
		if rune(nbr) < '0' || rune(nbr) > '9' {
			res = false
		}
	}
	return res

}

func Any(f func(string) bool, a []string) bool {
	res := false
	for _, word := range a {
		if f(word) == true {
			res = true
		}
	}
	return res
}
