// Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.
// These combinations are separated by a comma and a space.
package main

import "github.com/01-edu/z01"

func printComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; i <= '9'; i++ {
			l := j + 1
			for k := i; k <= '9'; k++ {
				for ; l <= '9'; l++ {
					z01.PrintRune((i))
					z01.PrintRune((j))
					z01.PrintRune(' ')
					z01.PrintRune((k))
					z01.PrintRune((l))

					if i < '9' || j < '8' || k < '9' || l < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				l = '0'
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	printComb2()
}
