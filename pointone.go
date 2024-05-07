// Write a function that takes a pointer to an int as argument and gives to this int the value of 1.
package main

import "fmt"

func pointOne(n *int) {
	*n = 1
}

func main() {
	num := 4
	pointOne(&num)
	fmt.Println(num)
}
