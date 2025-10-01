package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var CommandRegistry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(config) error
}

type config struct{}

func initCommands() {
	CommandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations on the map",
			callback:    commandMapb,
		},
	}
}

func commandHelp(c config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range CommandRegistry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandExit(c config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

var mapPage int

func commandMap(c config) error {
	mapPage++
	fetchMapPage()
	return nil
}

func commandMapb(c config) error {
	mapPage--
	fetchMapPage()
	return nil
}

func fetchMapPage() { //[]string {
	val, err := apiGet(fmt.Sprintf("/location-area?offset=%v&limit=20", (mapPage-1)*20))
	if err != nil {
		fmt.Printf("network error: %v\n", err)
		//return nil
	}

	var js LocationAreaResult
	if err := json.Unmarshal(val, &js); err != nil {
		fmt.Printf("failed to unmarshal response: %v\n", err)
		//return nil
	}

	//locations := []string{}
	for _, val := range js.Results {
		//locations = append(locations, val.Name)
		fmt.Println(val.Name)
	}
	//return locations
}
