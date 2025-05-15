package main

import (
	"fmt"
	"os"
	"sellsuki/security_research/command_execution/cmd"
)

func main() {
	fmt.Println("Command Execution Example")

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go [safe|unsafe] [command]")
		return
	}

	mode := os.Args[1]
	command := os.Args[2]

	switch mode {
	case "safe":
		err := cmd.SafeExecute(command)
		if err != nil {
			fmt.Printf("Safe execution error: %v\n", err)
		}
	case "unsafe":
		err := cmd.UnsafeExecute(command)
		if err != nil {
			fmt.Printf("Unsafe execution error: %v\n", err)
		}
	default:
		fmt.Println("Invalid mode, use 'safe' or 'unsafe'")
	}
}
