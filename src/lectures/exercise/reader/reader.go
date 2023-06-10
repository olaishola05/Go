package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	hello = "hello"
	bye   = "bye"
)

type Command string

func commandStats(cmdCounter, lineCounter int) {
	fmt.Printf("The commands stats\n numbers of command: %v\n number of non-blank line: %v\n", cmdCounter, lineCounter)
}

func messageResponder(c *Command, cmdCounter int) int {
	switch command := *c; command {
	case hello:
		fmt.Println("Command Response: Hi, Welcome to code lands, how can I help you?")
		cmdCounter += 1
	case bye:
		fmt.Println("Command Response: Thank you for stopping by.")
		cmdCounter += 1
	}

	return cmdCounter
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	lineCounter := 0
	commandCounter := 0

	for reader.Scan() {
		if strings.ToUpper(reader.Text()) == "Q" {
			break
		} else {
			command := Command(strings.ToLower(strings.TrimSpace(reader.Text())))
			commandCounter = messageResponder(&command, commandCounter)

			if command != "" {
				lineCounter += 1
			}
		}
	}

	commandStats(commandCounter, lineCounter)

}
