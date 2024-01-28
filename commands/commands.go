package commands

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	http "net/http"
	"io"
	"encoding/json"
)

type commandline struct {
	name string
	description string
	callback func() error
}

type maps struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
func commandMap() error {
	locationData := maps{}
	resp, err := http.Get("https://pokeapi.co/api/v2/location")
	if err != nil {
		fmt.Println("Unfortunately there was an error with the map command")
		return nil
	}
	
	body, err := io.ReadAll(resp.Body)
	
	resp.Body.Close()

	if resp.StatusCode > 299 {
		fmt.Printf("Something happened on our end %v", resp.StatusCode)
	}

	maperr := json.Unmarshal(body, &locationData)
	if maperr != nil {
		fmt.Printf("Error unmarshaling location data")
	}

	resp2, err2 := http.Get(locationData.Next)
	if err2 != nil {
		fmt.Println("Error fetching data")
	}

	body2, err := io.ReadAll(resp2.Body)
	resp2.Body.Close()


	fmt.Printf("%s", body2)

	// for _, mapName := range locationData.Results {
	// 	fmt.Printf("%s\n", mapName.Name)
	// }
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
			callback: commandMap,
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
		}
	}
}