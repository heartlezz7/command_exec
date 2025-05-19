# Command Execution Example (Go)

This project demonstrates **safe** and **unsafe** command execution patterns in Go.

## Usage

Run in safe mode (only allowed commands):

```bash
go run main.go safe date
```

Run in safe mode (only allowed commands):

```bash
 go run main.go unsafe 'touch test.go'
```

Run test nslookup execution (safe vs. unsafe):

Safe (no shell execution):

```bash
go run main.go nslookup "example.com; echo injected" safe
```

Unsafe (shell execution, possible injection):

```bash
go run main.go nslookup "example.com; echo injected" unsafe
```
