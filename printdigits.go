// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.
package main

import "github.com/01-edu/z01"

func printDigits() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func main() {
	printDigits()
}
