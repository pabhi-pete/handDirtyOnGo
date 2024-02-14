// types/pointers/begin/main.go
package main

func main() {
	// create a variable of type *T where T is an int
	var a *int
	// declare and assign `b` variable of type int
	b := 99
	// assign the address of b to a
	a = &b
	// print out the value of a which is the address of b
	println("Address of 'a' in hex is: ", a)
	// print out the value at the address of b
	println("The value of 'a' is: ", *a)

}
