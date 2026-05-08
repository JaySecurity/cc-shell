package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
prompt:
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = line[:len(line)-1]
		args := strings.Split(line, " ")
		command := args[0]
		switch command {
		case "exit":
			break prompt
		case "echo":
			echo(args[1:])
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func echo(args []string) {
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	fmt.Printf("\n")
}
