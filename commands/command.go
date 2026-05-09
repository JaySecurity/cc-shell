package command

import "fmt"

type Command struct {
	Name   string
	Action interface{}
	Type   string
}

func (c *Command) Register() {
	fmt.Println("Registered")

	testCmd := Command{
		Name: "echo",
		Type: "Builtin",
		Action: func(args []string) {
			fmt.Println()
		},
	}
	fmt.Println(testCmd.Name, "-", testCmd.Type)
}
