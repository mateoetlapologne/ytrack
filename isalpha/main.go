package main

import "fmt"

func main() {
	fmt.Print(isalpha(("s6a")))
}

func isalpha(str string) bool {
	res := true
	for index, letter := range str {
		_ = index
		if (letter < 'a' || letter > 'z') && (letter < 'A' || letter > 'Z') && (letter < '0' || letter > '9') {
			res = false
		}

	}
	return res
}
