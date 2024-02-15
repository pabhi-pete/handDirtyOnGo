// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	//
	// names := make([]string, 0)
	//alternate syntax
	names := []string{}
	// append 3 names to the slice
	names = append(names, "Jane")
	names = append(names, "Mary")
	names = append(names, "John")
	// print the slice
	// slice the slice using syntax slice[low:high]
	fmt.Println(names[:2]) // [Jane Mary]

	fmt.Println(names[1:])  // [Mary John]
	fmt.Println(names[1:3]) // [Mary John]
	fmt.Println(names[2:])  // [John]
	fmt.Println(names[:1])  // [Jane]
	// [John Jane Mary]
	fmt.Println(names[:3])
	fmt.Println(names[:])
}
