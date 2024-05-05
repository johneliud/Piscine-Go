// Write a program that prints the Latin alphabet in lowercase on a single line.
package main

import "github.com/01-edu/z01"

func printAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func main() {
	printAlphabet()
}