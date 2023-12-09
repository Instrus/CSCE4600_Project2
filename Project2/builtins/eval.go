package builtins

import (
	"fmt"
	"os/exec"
)

func Eval(args ...string) error {
	// Execute the command using the exec.Command function
	cmd := exec.Command(args[0], args[1:]...)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}

	// Print the output of the command (optional)
	fmt.Printf("Command output: %s\n", output)

	return nil
}
