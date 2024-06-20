package main

import (
	"bufio"
	"strings"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	for{
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			fmt.Println(strings.Split(input, "\n")[0] + ": command not found")
		}
	}
	
}
