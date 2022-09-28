// write a program that prints the name of the program.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
