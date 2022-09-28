package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", slicewithspace("Salut a tous c'est lasalle"))
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
