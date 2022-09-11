package main

import "fmt"

func main() {
	fmt.Println(printwordstables(slicewithspace("Hello how are you?")))
}

func slicewithspace(s string) []string {
	var res []string
	var timeres string
	for _, lettre := range s {
		if rune(lettre) == 32 || rune(lettre) == 9 || rune(lettre) == 10 || rune(lettre) == 27 {
			res = append(res, timeres)
			timeres = ""
		} else {
			timeres += string(lettre)
		}
	}
	res = append(res, timeres)
	return res

}

func printwordstables(carra []string) string {
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
