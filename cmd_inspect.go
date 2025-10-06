package main

import (
	"fmt"

	"github.com/crowspin/pokeapi"
)

func commandInspect(c *pokeapi.AppState) {
	var pkmnName string
	if c.Arg == "" {
		fmt.Println("Invalid use of command 'inspect'. See usage with command 'help'.")
		return
	} else {
		pkmnName = c.Arg
	}

	if val, ok := c.Pokedex[pkmnName]; ok {
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
		fmt.Println("You have not caught that pokemon yet...")
	}
}
