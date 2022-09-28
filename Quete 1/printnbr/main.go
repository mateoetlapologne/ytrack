package printnbr

import "github.com/01-edu/z01"

func main() {
	Printnbr(123)
}

func Printnbr(number int) {
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	if number > 9 {
		Printnbr(number / 10)
	}
	z01.PrintRune(rune(number%10 + 48))
}
