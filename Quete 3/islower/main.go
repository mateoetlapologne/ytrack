package main

import "fmt"

func main() {
	fmt.Println(isupper("salut"))
}

func isupper(str string) bool {
	res := true
	for index, lettre := range str {
		_ = index
		if lettre < 'a' || lettre > 'z' {
			res = false
		}
	}
	return res
}
