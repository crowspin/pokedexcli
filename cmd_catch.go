package main

import (
	"fmt"
	"math/rand"

	"github.com/crowspin/pokeapi"
)

func commandCatch(c *pokeapi.AppState) {
	var targetName, dest string
	if c.Arg == "" {
		fmt.Println("Invalid use of command 'catch'. See usage with command 'help'.")
		return
	} else {
		targetName = c.Arg
		dest = pokeapi.APIURL + "/pokemon/" + targetName
	}

	var js pokeapi.Pokemon
	if err := pokeapi.ApiGet(c, dest, &js); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", targetName)

	if rand.Intn(js.BaseExperience) < 25 {
		fmt.Printf("%s was caught!\n", targetName)
		c.Pokedex[targetName] = js
	} else {
		fmt.Printf("%s escaped!\n", targetName)
	}
}
