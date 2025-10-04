package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/crowspin/pokecache"
)

var CommandRegistry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

var cache pokecache.Cache

func initCommands() {
	cache = pokecache.NewCache(5 * time.Second)
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

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range CommandRegistry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *config) error {
	fetchMapPage(c, c.Next)
	return nil
}

func commandMapb(c *config) error {
	fetchMapPage(c, c.Previous)
	return nil
}

func fetchMapPage(c *config, dest string) {
	if dest == "" {
		dest = APIURL + "/location-area"
	}

	val, err := apiGet(dest)
	if err != nil {
		fmt.Printf("network error: %v\n", err)
	}

	var js LocationAreaResult
	if err := json.Unmarshal(val, &js); err != nil {
		fmt.Println(string(val))
		fmt.Printf("failed to unmarshal response: %v\n", err)
	}

	c.Next = js.Next
	c.Previous = js.Previous

	for _, val := range js.Results {
		fmt.Println(val.Name)
	}
}
