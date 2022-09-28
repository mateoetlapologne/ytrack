package main

import "fmt"

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	var tempchar string
	for i := len(s) - 1; i > -1; i-- {
		tempchar += string(s[i])
	}
	return tempchar

}
