package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Command struct {
	Name        string
	Args        []string
	CommandType string
}

func (c Command) Execute() {
	if c.Name == "exit" && len(c.Args) > 0 && c.Args[0] == "0" {
		os.Exit(0)
	} else if c.Name == "echo" {
		fmt.Println(strings.Join(c.Args, " "))
	} else if c.Name == "type" {
		subcommand := c.Args[0]
		if slices.Contains(builtinCommand, subcommand) {
			fmt.Println(subcommand + " is a shell builtin")
		} else {
			fmt.Println(subcommand + ": not found")
		}
	} else {
		fmt.Println(c.Name + ": command not found")
	}
}

var builtinCommand = []string{"cd", "pwd", "echo", "exit", "type"}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			input = strings.TrimSpace(input)
			slices := strings.Split(input, " ")
			command := CreateCommand(slices[0], slices[1:])
			command.Execute()
		}
	}
}

func CreateCommand(command string, args []string) Command {
	var commandType string = "unknown"
	if slices.Contains(builtinCommand, command) {
		commandType = "builtin"
	}
	return Command{
		Name:        command,
		Args:        args,
		CommandType: commandType,
	}
}
