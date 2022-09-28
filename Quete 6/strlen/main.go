package main

import "fmt"

func main() {
	l := strlen("Hello Word !")
	fmt.Print(l)
}

func strlen(s string) int {
	count := 0
	for _, c := range s {
		_ = c
		count++
	}
	return count
}
