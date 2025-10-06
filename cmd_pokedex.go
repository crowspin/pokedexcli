package main

import (
	"fmt"

	"github.com/crowspin/pokeapi"
)

func commandPokedex(c *pokeapi.AppState) {
	if len(c.Pokedex) > 0 {
		fmt.Println("Your Pokedex:")
		for _, p := range c.Pokedex {
			fmt.Printf(" - %s\n", p.Name)
		}
	} else {
		fmt.Println("You haven't caught any pokemon yet.")
	}
}
