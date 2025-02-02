// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["one"] = author{name: "First Person"}
	authors["two"] = author{name: "Second Persons"}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)
	anotherApproch()
	// read a value from the map with a known key
	fmt.Println(authors["one"])
	fmt.Println(authors["two"])
}

func anotherApproch() {
	authors := map[string]author{
		"one": {name: "Another Person"},
		"two": {name: "Next Persons"},
	}
	fmt.Println(authors)
}
