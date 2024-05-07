package main

import (
	"github.com/01-edu/z01"
)

func printNum(num int) {
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	digits := []int{}

	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}

func printNbr(num int) {
	printNum(num)
	z01.PrintRune('\n')
}

func main() {
	printNbr(23)
	printNbr(-123)
}
