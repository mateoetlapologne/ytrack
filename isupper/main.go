package main

import "fmt"

func main() {
	fmt.Println(isupper("SALUT"))
}

func isupper(str string) bool {
	res := true
	for index, lettre := range str {
		_ = index
		if lettre < 'A' || lettre > 'Z' {
			res = false
		}
	}
	return res
}
