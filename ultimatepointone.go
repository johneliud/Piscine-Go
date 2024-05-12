// Write a function that takes a pointer to a pointer to a pointer to an int as argument and gives to this int the value of 1.
package main

import "fmt"

func ultimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	ultimatePointOne(&n)
	fmt.Println(a)
}
