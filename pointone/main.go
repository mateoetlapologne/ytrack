package main

import (
	"fmt"
)

func main() {
	n := 0
	pointone(&n)
	fmt.Println(n)
}

func pointone(n *int) {
	*n = 1

}
