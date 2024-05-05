// Write a program that prints the Latin alphabet in lowercase in reverse order (from 'z' to 'a') on a single line.
package main

import "github.com/01-edu/z01"

func printReverseAlphabet() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func main() {
	printReverseAlphabet()
}
