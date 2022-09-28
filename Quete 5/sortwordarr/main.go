package main

import "fmt"

func main() {
	a := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3", "Y", "h", "6"}
	fmt.Print(sortwordarr(a))
}

func sortwordarr(a []string) []string {
	for i := 0; i+1 < len(a); i++ {
		n1 := a[i]
		n2 := a[i+1]
		if n1 > n2 {
			a[i] = n2
			a[i+1] = n1
			sortwordarr(a)
		}

	}
	return a
}
