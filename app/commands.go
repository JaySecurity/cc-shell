package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

func handleEcho(args []string, ctx *CommandContext) {
	args = args[1:]
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	fmt.Printf("\n")
}

func handleExit(_ []string, _ *CommandContext) {
	// fmt.Println("Goodbye!")
	os.Exit(0)
}

func pwd(_ []string, _ *CommandContext) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func getType(args []string, ctx *CommandContext) {
	if len(args) < 1 {
		return
	} else {
		cmd, ok := ctx.Registry[args[1]]
		if ok {
			fmt.Printf("%s is a shell %s\n", cmd.Name, cmd.Type)
		} else {
			filepath := checkCmd(args[1])
			if filepath != "" {
				fmt.Printf("%s is %s\n", args[1], filepath)
				return
			}
			fmt.Printf("%s: not found\n", args[1])
		}
	}
}

func checkCmd(cmd string) string {
	filepath, err := exec.LookPath(cmd)
	if err != nil {
		return ""
	}
	return filepath
}
