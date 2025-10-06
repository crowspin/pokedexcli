package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	Args     []string
	Pokedex  map[string]Pokemon
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
		"explore": {
			name:        "explore",
			description: "Lists the pokemon it's possible to encounter in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows the stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists Pokemon you've caught",
			callback:    commandPokedex,
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
		return
	}

	var js LocationAreaResult
	if err := json.Unmarshal(val, &js); err != nil {
		fmt.Printf("failed to unmarshal response: %v\n", err)
		return
	}

	c.Next = js.Next
	c.Previous = js.Previous

	for _, val := range js.Results {
		fmt.Println(val.Name)
	}
}

func commandExplore(c *config) error {
	val, err := apiGet(APIURL + "/location-area/" + c.Args[1])
	if err != nil {
		fmt.Printf("network error: %v\n", err)
		return nil
	}

	var js LocationAreaData
	if err := json.Unmarshal(val, &js); err != nil {
		fmt.Printf("failed to unmarshal response: %v\n", err)
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, pkmn := range js.PokemonEncounters {
		fmt.Println(" -", pkmn.Pokemon.Name)
	}
	return nil
	//I can't wait to finish this project so I can refactor it into a reasonable application. I'm never going to use it but this is being presented in such an overcomplicated manner.
	//What may be lost in translation here is that I'm trying to finish the rest of the Python/Go path before the end of the month. I'm not sure how that's going to go, but I just don't have time to fully reorganize myself every single lesson.
}

func commandCatch(c *config) error {
	targetName := c.Args[1]
	val, err := apiGet(APIURL + "/pokemon/" + targetName)
	if err != nil {
		fmt.Printf("network error: %v\n", err)
		return nil
	}

	var js Pokemon
	if err := json.Unmarshal(val, &js); err != nil {
		fmt.Printf("failed to unmarshal response: %v\n", err)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", targetName)

	if rand.Intn(js.BaseExperience) < 25 {
		fmt.Printf("%s was caught!\n", targetName)
		c.Pokedex[targetName] = js
	} else {
		fmt.Printf("%s escaped!\n", targetName)
	}
	return nil
	//handle no args[1]
}

func commandInspect(c *config) error {
	if val, ok := c.Pokedex[c.Args[1]]; ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Printf("Stats:\n")
		for _, v := range val.Stats {
			fmt.Printf("  -%s: %v\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, v := range val.Types {
			fmt.Printf("  - %s\n", v.Type.Name)
		}
	} else {
		fmt.Printf("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(c *config) error {
	fmt.Println("Your Pokedex:")
	for _, p := range c.Pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
