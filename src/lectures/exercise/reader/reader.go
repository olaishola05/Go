//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	hello = iota
	bye
)

type Message string

func messageResponder(msg string) string {
	newMsg := Message(msg)
	num := 0
	switch num {
	case hello:
		return "Hi, Welcome to code lands"
	case bye:
		return "Thank you for stopping by."
	}

	return string(newMsg)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	lineCounter, commands := 0, 0

	for {
		command, commandErr := reader.ReadString(' ')
		n := strings.TrimSpace(command)
		message := messageResponder(n)
		if n != "" {
			lineCounter += 1
			commands += 1
			fmt.Println(message)
		}

		if commandErr == io.EOF {
			break
		}

		if commandErr != nil {
			fmt.Println("error reading Stdin:", commandErr)
		}
	}
	fmt.Printf("total Stats: \nNumber non-blank lines: %v\nNumbers of commands entered: %v\n", lineCounter, commands)
}
