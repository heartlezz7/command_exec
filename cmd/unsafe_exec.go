package cmd

import (
	"fmt"
	"os/exec"
)

// Unsafe execution (direct from user input — dangerous)
func UnsafeExecute(userInput string) error {
	// Directly pass user input to shell (DANGEROUS!)
	cmd := exec.Command("bash", "-c", userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command error: %w", err)
	}

	fmt.Println("Output:")
	fmt.Println(string(output))
	return nil
}
