package main

import (
	"bytes"
	"errors"
)

type Redirect int

const (
	None Redirect = iota
	Write
	Append
	Pipe
)

func CleanArgs(input []byte) ([]byte, error) {
	var output []byte
	args := bytes.NewBuffer(output)

	isSingleQuote := false
	isDoubleQuote := false
	escaped := false

	for i, b := range input {

		if escaped {
			args.WriteByte(b)
			escaped = false
			continue
		}
		if b == '\\' && !isSingleQuote {
			escaped = true
			continue
		}
		if b == '"' && !isSingleQuote {
			isDoubleQuote = !isDoubleQuote
			continue
		}
		if b == '\'' && !isDoubleQuote {
			isSingleQuote = !isSingleQuote
			continue
		}

		if b == ' ' && !isSingleQuote && !isDoubleQuote && input[i-1] == ' ' {
			continue
		}
		args.WriteByte(b)
	}
	return args.Bytes(), nil
}

func ParseCmd(input []byte) (*CommandInput, error) {
	var command string
	var output []byte
	redirect := None
	destination := ""

	for i, val := range input {
		if val == ' ' {
			command = string(input[:i])
			if len(input) > i+1 {
				input = input[i+1:]
			}
			break
		}
	}
	if command == "" {
		command = string(input)
		input = nil
	}
	output = input
	for i, val := range input {
		if string(input[i:i+2]) == "1>" {
			redirect = Write
			if i+2 > len(input)-1 {
				return nil, errors.New("invalid command")
			}
			destination = string(bytes.TrimSpace(input[i+2:]))
			output = bytes.TrimSpace(input[:i])
			break
		}
		if val == '>' {
			redirect = Write
			if i+1 > len(input)-1 {
				return nil, errors.New("invalid command")
			}
			destination = string(bytes.TrimSpace(input[i+1:]))
			output = bytes.TrimSpace(input[:i])
			break
		}
	}
	return &CommandInput{
		command:  command,
		args:     output,
		redirect: redirect,
		dest:     destination,
	}, nil
}
