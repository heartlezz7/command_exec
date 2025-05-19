package main

import (
	"fmt"
	"os"
	"os/exec"
	"sellsuki/security_research/command_execution/cmd"
)

func main() {
	fmt.Println("Command Execution Example")

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go [safe|unsafe] [command]")
		fmt.Println("  go run main.go nslookup [domain] [safe|unsafe]")
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
	case "nslookup":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go nslookup [domain] [safe|unsafe]")
			return
		}
		domain := os.Args[2]
		mode := os.Args[3]

		var out []byte
		var err error

		if mode == "safe" {
			out, err = exec.Command("nslookup", domain).CombinedOutput()
		} else if mode == "unsafe" {
			out, err = exec.Command("bash", "-c", "nslookup "+domain).CombinedOutput()
		} else {
			fmt.Println("Invalid nslookup mode, use 'safe' or 'unsafe'")
			return
		}

		if err != nil {
			fmt.Printf("Error running nslookup: %v\n", err)
		}
		fmt.Println(string(out))
	default:
		fmt.Println("Invalid mode, use 'safe' or 'unsafe'")
	}
}
