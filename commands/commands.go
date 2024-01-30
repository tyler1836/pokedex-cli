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
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		} `json:"results"`
}

var url string = "https://pokeapi.co/api/v2/location"
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
	// pagination is working however value of url needs to be updated somehow globally. As is when mapb is called after map the next 20 show and the prev 20 show when map is called after mapb
	locationData := maps{}
	resp, err := http.Get(url)
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
	
	
	for _, mapName := range locationData.Results {
		fmt.Printf("%s\n", mapName.Name)
	}
	url = locationData.Next
	return nil
}
func commandMapb() error {
	locationData := maps{}
	resp, err := http.Get(url)
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
	if locationData.Previous != "" {
		url = locationData.Previous
	}

	for _, mapName := range locationData.Results {
			fmt.Printf("%s\n", mapName.Name)
	}
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
		"mapb": commandline{
			name: "mapb",
			description: "Display previous locations from the Pokemon World",
			callback: commandMapb,
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