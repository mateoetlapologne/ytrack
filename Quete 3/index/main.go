package main

import "fmt"

func main() {
	fmt.Println(index("Hello", "lo"))
}

func index(fstr string, sstring string) int {
	x := len(sstring)
	for i := 0; x <= len(fstr); {
		if fstr[i:x] == sstring {
			return i
		} else {
			i++
			x++
		}
	}
	return -1

}
