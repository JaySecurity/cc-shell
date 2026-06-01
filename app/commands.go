package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func handleChange(input *CommandInput, _ *CommandContext) (output []byte, err error) {
	home := os.Getenv("HOME")
	args := string(input.args)
	if len(args) <= 1 {
		os.Chdir(home)
	} else if path, found := strings.CutPrefix(args, "~"); found {
		path = filepath.Join(home, path)
		if err := os.Chdir(path); err != nil {
			output = fmt.Appendf(nil, "cd: %s: No such file or directory\n", path)
		}
	} else if path, found := strings.CutPrefix(args, "./"); found {
		if err := os.Chdir(path); err != nil {
			output = fmt.Appendf(nil, "cd: %s: No such file or directory\n", path)
		}
	} else if strings.HasPrefix(args, "../") {
		if err := os.Chdir(args); err != nil {
			output = fmt.Appendf(nil, "cd: %s: No such file or directory\n", args)
		}
	} else if strings.HasPrefix(args, "/") {
		if err := os.Chdir(args); err != nil {
			output = fmt.Appendf(nil, "cd: %s: No such file or directory\n", args)
		}
	} else {
		if err := os.Chdir(args); err != nil {
			output = fmt.Appendf(nil, "cd: %s: No such file or directory\n", args)
		}
	}
	return output, nil
}

func handleEcho(input *CommandInput, ctx *CommandContext) (output []byte, err error) {
	args := input.args
	output = append(args, '\n')
	return output, nil
}

func handleExit(_ *CommandInput, _ *CommandContext) ([]byte, error) {
	return []byte("quit"), nil
}

func pwd(_ *CommandInput, _ *CommandContext) (output []byte, err error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	output = fmt.Appendf([]byte(dir), "\n")
	return output, nil
}

func getType(input *CommandInput, ctx *CommandContext) (output []byte, err error) {
	command := string(input.args)
	if len(command) < 1 {
		return
	} else {
		cmd, ok := ctx.Registry[command]
		if ok {
			output = fmt.Appendf(nil, "%s is a shell %s\n", cmd.Name, cmd.Type)
		} else {
			filepath := checkCmd(command)
			if filepath != "" {
				output = fmt.Appendf(nil, "%s is %s\n", command, filepath)
				return output, nil
			}
			output = fmt.Appendf(nil, "%s: not found\n", command)
		}
	}
	return output, err
}

func checkCmd(cmd string) string {
	filepath, err := exec.LookPath(cmd)
	if err != nil {
		return ""
	}
	return filepath
}
