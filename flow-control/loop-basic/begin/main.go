// flow-control/loop-basic/begin/main.go
package main

import (
	"fmt"
)

func main() {
	// declare a string to iterate over
	l := "automation"

	// iterate over the string with basic for loop
	for i := 0; i < len(l); i++ {
		fmt.Println(i, ":", string(l[i]))
	}
}
