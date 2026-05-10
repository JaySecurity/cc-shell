package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func handleChange(args []string, _ *CommandContext) {
	home := os.Getenv("HOME")
	if len(args) <= 1 {
		os.Chdir(home)
	} else if path, found := strings.CutPrefix(args[1], "~"); found {
		path = filepath.Join(home, path)
		if err := os.Chdir(path); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", path)
		}
	} else if path, found := strings.CutPrefix(args[1], "./"); found {
		if err := os.Chdir(path); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", path)
		}
	} else if strings.HasPrefix(args[1], "../") {
		if err := os.Chdir(args[1]); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", args[1])
		}
		// cwd, err := os.Getwd()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// segments := strings.Split(args[1], "/")
	} else if strings.HasPrefix(args[1], "/") {
		if err := os.Chdir(args[1]); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", args[1])
		}
	} else {
		if err := os.Chdir(args[1]); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", args[1])
		}
	}
}
