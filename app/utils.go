package main

import (
	"fmt"
	"strings"
)

func ParseArgs(input string) ([]string, error) {
	var args []string
	var current strings.Builder

	isSingleQuote := false
	isDoubleQuote := false
	// escaped := false

	for _, r := range input {

		if r == '"' && !isSingleQuote {
			isDoubleQuote = !isDoubleQuote
			continue
		}
		if r == '\'' && !isDoubleQuote {
			isSingleQuote = !isSingleQuote
			continue
		}

		if r == ' ' && !isSingleQuote && !isDoubleQuote {
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
			continue
		}
		current.WriteRune(r)
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	fmt.Println(args, len(args))
	return args, nil
}
