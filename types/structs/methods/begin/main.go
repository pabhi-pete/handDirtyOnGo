// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
//
func authorFullName(first string, last string) {
	fmt.Println("The author's name is: ", first, last)
}

//alternate way to get author fullname
func (a author) getFullName() string {
	return a.first + " " + a.last
}

func main() {
	// initialize author
	a := author{
		first: "Marcus",
		last:  "Aurelius",
	}

	// print the author's full name
	authorFullName(a.first, a.last)
	fmt.Println(a.getFullName())
}
