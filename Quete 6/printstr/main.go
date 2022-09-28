package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {
	for _, chara := range s {
		z01.PrintRune(chara)

	}

}
