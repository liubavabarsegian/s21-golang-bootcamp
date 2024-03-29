package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./myXargs <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)
	var args []string

	for scanner.Scan() {
		args = append(args, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "executing command:", err)
		os.Exit(1)
	}

}
