package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

// Safe execution (allowlist commands only)
func SafeExecute(userInput string) error {
	// Allow only these commands
	allowedCommands := map[string]string{
		"date":   "date",
		"uptime": "uptime",
	}

	cmdName, ok := allowedCommands[userInput]
	if !ok {
		return fmt.Errorf("command not allowed")
	}

	cmd := exec.Command(cmdName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command error: %w", err)
	}

	fmt.Println("Output:")
	fmt.Println(strings.TrimSpace(string(output)))
	return nil
}
