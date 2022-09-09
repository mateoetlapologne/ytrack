package main

import "fmt"

func main() {
	fmt.Print(concat("Hello", "Salut"))
}

func concat(str1 string, str2 string) string {
	return str1 + str2
}
