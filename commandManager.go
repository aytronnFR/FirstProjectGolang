package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	name string
	Execute func(args []string) bool
	usage string
}

var commands []Command

func commandManager() {
	defer commandQueue.Done()
	var args []string = consoleRead()

	inputCommand(args)
}

func registerCommand(command Command) {
	commands = append(commands, command)
}

func inputCommand(args []string) {

	command := args[0]

	if command == "help" {
		showHelp()
		return
	} else if command == "stop" {
		go shutdown()
		return
	}

	for _, command1 := range commands {
		if command1.name == command {
			if !command1.Execute(args) {
				info("Error while executing the command!")
			}
			return
		}
	}
	info("Command doesn't exist !")
}

func showHelp()  {
	fmt.Printf(InfoColor, "Help commands: ")
	fmt.Printf(InfoColor, "stop : allow to stop the application")
	for _, command := range commands {
		help := command.usage
		if help != "" {
			fmt.Printf(InfoColor, help)
		}
	}
}

func consoleRead() []string {
	input := bufio.NewReader(os.Stdin)
	text, err := input.ReadString('\n')
	failOnError(err, "Error when send message")
	text = strings.ReplaceAll(text, "\n", "")
	return strings.Split(text, " ")
}
