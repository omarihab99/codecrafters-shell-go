package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			input = strings.TrimSpace(input)
			slices := strings.Split(input, " ")
			handleCommand(slices[0], slices[1:])
		}
	}
}
// handleCommand is a function that takes a command and its arguments as input.
//
// Parameters:
// - command: a string representing the command to be executed.
// - args: a slice of strings representing the arguments passed to the command.
//
// Return type: None.
func handleCommand(command string, args []string) {
	if command == "exit" && len(args) > 0 && args[0] == "0" {
		os.Exit(0)
	} else {
		fmt.Println(command + ": command not found")
	}
}
