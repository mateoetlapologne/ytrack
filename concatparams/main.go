package main

import (
	"fmt"
)

func main() {
	test := []string{"Hello", "Caca", "Prout"}
	fmt.Println(concatparams(test))
}

func concatparams(carra []string) string {
	var res string
	for index, word := range carra {
		if index+1 != len(carra) {
			res += word
			res += "\n"
		} else {
			res += word
		}
	}
	return res

}
