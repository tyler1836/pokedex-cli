package commands

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	maps "github.com/tyler1836/pokedex-cli/pokecall"
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
		"map": commandline{
			name: "map",
			description: "Display locations from the Pokemon World",
			callback: maps.CommandMap,
		},
		"mapb": commandline{
			name: "mapb",
			description: "Display previous locations from the Pokemon World",
			callback: maps.CommandMapb,
		},
	}

	reader := bufio.NewReader(os.Stdin) 
	
	for {
		fmt.Println("Pokedex <")
		input, _ := reader.ReadString('\n')
		// trim whitespace and make lowercase for switch
		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
			case "help":
				commands["help"].callback()
			case "exit":
				commands["exit"].callback() 
			case "map":
				commands["map"].callback()
			case "mapb":
				commands["mapb"].callback()
		}
	}
}