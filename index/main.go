package main

import "fmt"

func main() {
	fmt.Println(index("Hello", "lo"))
}

func index(fstr string, sstring string) int {
	for index, letter := range fstr {
		_ = letter
		_ = index
		if string(letter) == sstring {
			return index
		}
	}
	return -1
}
