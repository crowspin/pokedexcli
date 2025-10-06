package main

import (
	"fmt"

	"github.com/crowspin/pokeapi"
)

func commandHelp(c *pokeapi.AppState) {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range CommandRegistry {
		fmt.Printf("%s: %s\n", v.usage, v.description)
	}
}
