package main

import (
	"coursecontent/src/lectures/demo/pkg/display"
	"coursecontent/src/lectures/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("An exciting message")
}
