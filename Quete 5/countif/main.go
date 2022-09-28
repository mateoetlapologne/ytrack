package main

import "fmt"

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)

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

func CountIf(f func(string) bool, a []string) int {
	res := 0
	for _, str := range a {
		if f(str) == true {
			res++
		}
	}
	return res

}
