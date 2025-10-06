package main

import "github.com/crowspin/pokeapi"

type replCommand struct {
	usage       string
	description string
	callback    func(*pokeapi.AppState)
}

var CommandRegistry map[string]replCommand

func initCommands() {
	CommandRegistry = map[string]replCommand{
		"exit": {
			usage:       "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			usage:       "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			usage:       "map",
			description: "Displays the next page of locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			usage:       "mapb",
			description: "Displays the previous page of locations on the map",
			callback:    commandMapb,
		},
		"explore": {
			usage:       "explore <locationName>",
			description: "Lists the pokemon it's possible to encounter in an area",
			callback:    commandExplore,
		},
		"catch": {
			usage:       "catch <pokemonName>",
			description: "Tries to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			usage:       "inspect <pokemonName>",
			description: "Shows the stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			usage:       "pokedex",
			description: "Lists Pokemon you've caught",
			callback:    commandPokedex,
		},
	}
}
