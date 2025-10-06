package main

import (
	"fmt"

	"github.com/crowspin/pokeapi"
)

func commandExplore(c *pokeapi.AppState) {
	var dest string
	if c.Arg == "" {
		fmt.Println("Invalid use of command 'explore'. See usage with command 'help'.")
		return
	} else {
		dest = pokeapi.APIURL + "/location-area/" + c.Arg
	}

	var js pokeapi.LocationAreaData
	if err := pokeapi.ApiGet(c, dest, &js); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Found Pokemon:")
	for _, pkmn := range js.PokemonEncounters {
		fmt.Println(" -", pkmn.Pokemon.Name)
	}
}
