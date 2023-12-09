package builtins

import (
	"fmt"
)

func Echo(args ...string) error {

	// create a new slice to hold the elements
	elements := make([]interface{}, len(args))

	// add each element to an array to display back to user
	for i, arg := range args {
		elements[i] = arg
	}

	// print the elements
	fmt.Println(elements...)
	return nil
}
