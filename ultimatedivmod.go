/*
Create a function ultimateDivMod should divide the dereferenced value of a by the dereferenced value of b.
Store the result of the division in the int which a points to.
Store the remainder of the division in the int which b points to.
*/
package main

import "fmt"

func ultimateDivMod(a *int, b *int) {
	c := *a
	*a = *a / *b
	*b = c % *b
}

func main() {
	a := 13
	b := 2
	ultimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
