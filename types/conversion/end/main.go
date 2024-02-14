package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	a := 42.5
	b := uint(a)
	
	x := 100.203
	y := uint(x)

	//execute the value and indicate the type of variable
	fmt.Printf("%v=%T, %v=%T\n", a, a, b, b)
	fmt.Printf("%v=%T, %v=%T\n", x, x, y, y)
}
