package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodBye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numLines := 0
	numCommands := 0

	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())
			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("comand responce: hi")
			case CmdGoodBye:
				numCommands += 1
				fmt.Println("command responce: bye")
			}
			if text != "" {
				numLines += 1
			}
		}
	}

	fmt.Printf("you have ntered %v lines \n", numLines)
	fmt.Printf("you have entered %v lines \n", numCommands)
}
