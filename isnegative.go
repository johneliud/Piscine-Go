// Write a function that prints 'T' (true) on a single line if the int passed as parameter is negative, otherwise it prints 'F' (false).
package main

import "github.com/01-edu/z01"

func isNegative(num int) {
	if num > 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')
}

func main() {
	isNegative(23)
}
