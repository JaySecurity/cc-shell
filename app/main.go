package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	ctx := &CommandContext{Registry: commands}
	fmt.Println("Shell started. Type 'exit' to quit.")
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = line[:len(line)-1]
		args := strings.Split(line, " ")
		command, ok := commands[args[0]]
		if ok {
			command.Action(args, ctx)
		} else {
			fmt.Printf("%s: command not found\n", args[0])
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

func getType(args []string, ctx *CommandContext) {
	if len(args) < 1 {
		return
	} else {
		cmd, ok := ctx.Registry[args[1]]
		if ok {
			fmt.Printf("%s is a shell %s", cmd.Name, cmd.Type)
		} else {
			fmt.Printf("%s: not found\n", args[1])
		}
	}
	fmt.Printf("\n")
}

func register(commands map[string]Command, name string, cmdType string, action CommandHandler) {
	fmt.Println("Registered")

	cmd := Command{
		Name:   name,
		Type:   cmdType,
		Action: action,
	}

	commands[name] = cmd
}
