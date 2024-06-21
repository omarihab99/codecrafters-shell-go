package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type Command struct {
	Name        string
	Args        []string
	CommandType string
}

func (c Command) Execute() {
	if c.Name == "exit" && len(c.Args) > 0 && c.Args[0] == "0" { // exit
		os.Exit(0)
	} else if c.Name == "echo" { // echo
		fmt.Println(strings.Join(c.Args, " "))
		return
	} else if c.Name == "type" { // type
		fmt.Println(c.CommandType)
		return
	} else if c.Name == "pwd" { //pwd
		wd, _ := os.Getwd()
		fmt.Println(wd)
		return
	} else if c.Name == "cd" {
		path := c.Args[0]
		if path == "~" {
			path = os.Getenv("HOME")
		}
		err := os.Chdir(path)
		if err != nil {
			fmt.Println("cd: " + path + ": No such file or directory")
		}
		return
	}
	fp := searchPath(c.Name)
	if fp != "" { // executable
		cmd := exec.Command(fp, c.Args...)
		output, err := cmd.CombinedOutput()
		if err == nil {
			fmt.Println(strings.TrimSpace(string(output)))
			return
		}
	}
	fmt.Println(c.Name + ": command not found") // command not found
}
func handleType(command string, Args []string) string {
	subcommand := command
	if len(Args) > 0 {
		subcommand = Args[0]
	}
	if slices.Contains(builtinCommand, subcommand) {
		return subcommand + " is a shell builtin"
	}
	fp := searchPath(subcommand)
	if fp != "" {
		return fp
	}
	return subcommand + ": not found"
}
func searchPath(subcommand string) string {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		filePath := filepath.Join(path, subcommand)
		if _, err := os.Stat(filePath); err == nil {
			return filePath
		}
	}
	return ""
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
	commandType := handleType(command, args)
	return Command{
		Name:        command,
		Args:        args,
		CommandType: commandType,
	}
}
