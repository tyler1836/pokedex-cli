package pokecall

import (
	"fmt"
	http "net/http"
	"io"
	"encoding/json"
	cache "github.com/tyler1836/pokedex-cli/pokecache"
)

type maps struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		} `json:"results"`
}

var urlNext string = "https://pokeapi.co/api/v2/location"
var urlPrev string = "https://pokeapi.co/api/v2/location"

func CommandMap() error {
	locationData := maps{}
	resp, err := http.Get(urlNext)
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
	urlNext = locationData.Next
	urlPrev = locationData.Previous
	return nil
}
func CommandMapb() error {
	locationData := maps{}
	resp, err := http.Get(urlPrev)
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
		urlNext = locationData.Next
		urlPrev = locationData.Previous
	}

	for _, mapName := range locationData.Results {
			fmt.Printf("%s\n", mapName.Name)
	}
	return nil
}