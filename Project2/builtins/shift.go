package builtins

import "fmt"

func Shift(args ...string) error {
	if len(args) == 0 {
		return nil
	}

	// Skip the first item and print the remaining items
	for _, arg := range args[1:] {
		fmt.Println(arg)
	}

	return nil
}
