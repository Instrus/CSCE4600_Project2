package builtins

import (
	"fmt"
	"os/exec"
)

func Eval(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}
	fmt.Printf("Command output: %s\n", output)
	return nil
}
