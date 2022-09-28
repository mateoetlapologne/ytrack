package main

import "fmt"

func main() {
	fmt.Print(capitalyze("Heo !op!"))
}

func capitalyze(str string) string {
	word := ""

	for i, letter := range str {
		if i == 0 && letter >= 'a' && letter <= 'z' && (str[i] < 'A' || str[i] > 'Z') && (str[i] < '0' || str[i] > '9') {
			i++
			letter += -32
			word += string(letter)
		} else if letter >= 'a' && letter <= 'z' && (str[i-1] < 'A' || str[i-1] > 'Z') && (str[i-1] < '0' || str[i-1] > '9') && (str[i-1] >= 32 && str[i-1] <= 47) {
			i++
			letter += -32
			word += string(letter)
		} else {
			word += string(letter)
			i++
		}

	}
	return word
}
