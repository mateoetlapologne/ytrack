package main

import "fmt"

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 1, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func f(a int, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func IsSorted(f func(a int, b int) int, a []int) bool {
	res := true
	for index, nbr := range a {
		if index < len(a)-1 {
			if f(nbr, a[index+1]) == 1 {
				res = false
			}

		}
	}
	return res
}
