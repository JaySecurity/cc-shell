package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type CommandContext struct {
	Registry map[string]Command
}

type CommandHandler func(input *CommandInput, ctx *CommandContext) ([]byte, error)

type Command struct {
	Name   string
	Action CommandHandler
	Type   string
}
type CommandInput struct {
	command  string
	args     []byte
	redirect Redirect
	dest     string
}

func main() {
	commands := make(map[string]Command)
	register(commands, "echo", "builtin", handleEcho)
	register(commands, "exit", "builtin", handleExit)
	register(commands, "type", "builtin", getType)
	register(commands, "pwd", "builtin", pwd)
	register(commands, "cd", "builtin", handleChange)

	ctx := &CommandContext{Registry: commands}

	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = line[:len(line)-1]
		output, err := CleanArgs(line)
		if err != nil {
			log.Fatal(err)
		}

		input, err := ParseCmd(output)
		if err != nil {
			fmt.Println(err)
		}
		var out []byte
		command, ok := commands[input.command]
		if ok {
			out, err = command.Action(input, ctx)
		} else if checkCmd(input.command) != "" {
			if len(input.args) > 0 {
				args := strings.Fields(string(input.args))
				out, err = exec.Command(input.command, args...).Output()
			} else {
				out, err = exec.Command(input.command).Output()
			}
		} else {
			out = fmt.Appendf(nil, "%s: command not found\n", input.command)
		}
		if err != nil {
			if exitError, ok := errors.AsType[*exec.ExitError](err); ok {
				fmt.Printf("%s\n", exitError.Stderr)
			}
		}
		if string(out) == "quit" {
			os.Exit(0)
		}
		switch input.redirect {
		case None:
			fmt.Printf("%s", string(out))
		case Write:
			// fmt.Println(input.redirect, input.dest)
			err := os.WriteFile(input.dest, out, 0o666)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func register(commands map[string]Command, name string, cmdType string, action CommandHandler) {
	cmd := Command{
		Name:   name,
		Type:   cmdType,
		Action: action,
	}
	commands[name] = cmd
}
