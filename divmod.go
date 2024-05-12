/*
Write a function that will be as below.
This function will divide the int a and b.
The result of this division will be stored in the int pointed by div.
The remainder of this division will be stored in the int pointed by mod.
*/
package main

import "fmt"

func divMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	divMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
