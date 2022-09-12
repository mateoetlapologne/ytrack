package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func IsPrime(a int) bool {
	res := true
	for i := a - 1; i > 1; i-- {
		if a%i == 0 {
			res = false
		}
	}
	return res
}

func Map(f func(int) bool, a []int) []bool {
	var fres []bool
	for _, nbr := range a {
		if f(nbr) == true {
			fres = append(fres, true)
		} else {
			fres = append(fres, false)
		}
	}
	return fres
}
