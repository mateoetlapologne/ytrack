package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func Swap(s *int, ss *int) {
	tempres := *ss
	*ss = *s
	*s = tempres

}
