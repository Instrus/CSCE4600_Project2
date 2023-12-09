package builtins

import (
	"fmt"
	"os"
)

func PrintWorkingDirectory() error {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("Path\n----\n" + currentDir + "\n")
	return nil
}
