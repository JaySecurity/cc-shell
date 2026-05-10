package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type CommandContext struct {
	Registry map[string]Command
}

type CommandHandler func(args []string, ctx *CommandContext)

type Command struct {
	Name   string
	Action CommandHandler
	Type   string
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
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = line[:len(line)-1]
		// args, err := shlex.Split(line)
		args, err := ParseArgs(line)
		if err != nil {
			log.Fatal(err)
		}
		command, ok := commands[args[0]]
		if ok {
			command.Action(args, ctx)
		} else if checkCmd(args[0]) != "" {
			out, err := exec.Command(args[0], args[1:]...).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s: command not found\n", args[0])
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
