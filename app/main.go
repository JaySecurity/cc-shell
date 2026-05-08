package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: command not found\n", line[:len(line)-1])
	}
}
