package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%#v\n", slicewithspace("HelloHAhowHAareHAyou?", "HA"))
}

func slicewithspace(s string, c string) []string {
	var res []string
	var tempres string
	x := len(c)
	lastch := s[len(s)-1]
	for i := 0; x < len(s)+1; {
		if s[i:x] == c {
			res = append(res, tempres)
			tempres = ""
			i++
			x++
		} else {
			tempres += string(s[i])
			i++
			x++
		}
	}
	tempres += string(lastch)
	res = append(res, tempres)
	var finalres []string
	var newword string
	for index, word := range res {
		c5 := len(word)
		if index != 0 {
			newword = word[1:c5]
			finalres = append(finalres, newword)
		} else {
			newword = word
			finalres = append(finalres, newword)
		}

	}
	return finalres
	strconv.AppendBool()
}
