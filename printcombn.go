// Write a function that prints all possible combinations of n different digits in ascending order.
// n will be defined as : 0 < n < 10
// Below are the references for the printing format expected.
// (for n = 1) '0, 1, 2, 3, ..., 8, 9'
// (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'
package main

import "github.com/01-edu/z01"

func printCombinations(digits []rune, start, index int) {
	if index == len(digits) {
		printCombination(digits)
		z01.PrintRune(',')
		z01.PrintRune(' ')
		return
	}
	for i := start; i <= 9; i++ {
		digits[index] = rune(i + '0')
		printCombinations(digits, i+1, index+1)
	}
}

func printCombination(digits []rune) {
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func printCombN(n int) {
	if n <= 0 || n > 9 {
		return
	}
	digits := make([]rune, n)
	printCombinations(digits, 0, 0)
}

func main() {
	printCombN(3)
}
