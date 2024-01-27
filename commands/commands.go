package commands

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type commandline struct {
	name string
	description string
	callback func() error
}

var commands map[string]commandline

func commandHelp() error {
	fmt.Printf("Available commands:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s:\n %s\n", cmd.name, cmd.description)
	}
	return nil
}
func commandExit() error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return nil
}

func StartPokedex() {
	commands = map[string]commandline{
		"help": commandline{
			name: "help",
			description: "Displays help message",
			callback: commandHelp,
		},
		"exit": commandline{
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
	}

	reader := bufio.NewReader(os.Stdin) 
	
	for {
		fmt.Println("Pokedex <")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
			case "help":
				commands["help"].callback()
			case "exit":
				commands["exit"].callback() 
		}
	}
}