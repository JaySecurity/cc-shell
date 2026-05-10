package main

import (
	"fmt"
	"os"
	"strings"
)

func handleChange(args []string, _ *CommandContext) {
	path := ""
	home := os.Getenv("HOME")
	if len(args) <= 1 {
		os.Chdir(home)
	} else if strings.HasPrefix(args[1], "~") {
		fmt.Println("Relative to Home")
	} else if strings.HasPrefix(args[1], "./") {
		fmt.Println("Relative to CWD")
	} else if strings.HasPrefix(args[1], "../") {
		fmt.Println("Back from CWD ")
	} else if strings.HasPrefix(args[1], "/") {
		fmt.Println("Absolute")
		if err := os.Chdir(args[1]); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", args[1])
		}
	} else {
		fmt.Println("Relative to CWD")
	}
	fmt.Printf("Change to %s", path)
}
