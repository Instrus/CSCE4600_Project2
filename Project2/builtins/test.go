package builtins

import (
	"fmt"
	"os"
)

func Test(args ...string) error {

	var arguments []string

	//apprend each args to argument
	for _, arg := range args {
		arguments = append(arguments, arg)
	}

	// stop function if less than 2 arguments were passed
	if len(arguments) < 2 {
		fmt.Println("Error: please enter 2 arguments")
		return nil
	}
	// stop the function if the number of arguments is greater than 2
	if len(arguments) > 2 {
		fmt.Println("Error: cannot handle more than 2 arguments")
		return nil
	}

	if arguments[0] == arguments[1] {
		fmt.Printf("%s and %s are equal.\n", arguments[0], arguments[1])
	} else {
		fmt.Printf("%s and %s are not equal.\n", arguments[0], arguments[1])
	}

	// test for inequality
	if arguments[0] > arguments[1] {
		fmt.Printf("%s is greater than %s.\n", arguments[0], arguments[1])
	} else {
		fmt.Printf("%s is less than %s.\n", arguments[0], arguments[1])
	}

	// test if file exists
	if _, err := os.Stat(arguments[0]); err == nil {
		fmt.Printf("The file \"%s\" exists.\n", arguments[0])
	} else {
		fmt.Printf("The file \"%s\" does not exist.\n", arguments[0])
	}

	return nil

}
