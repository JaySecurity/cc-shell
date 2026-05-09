package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	register(commands, "pwd", "builtin", pwd)

	ctx := &CommandContext{Registry: commands}

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

func register(commands map[string]Command, name string, cmdType string, action CommandHandler) {
	cmd := Command{
		Name:   name,
		Type:   cmdType,
		Action: action,
	}
	commands[name] = cmd
}
