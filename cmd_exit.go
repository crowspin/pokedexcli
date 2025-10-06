package main

import (
	"fmt"
	"os"

	"github.com/crowspin/pokeapi"
)

func commandExit(c *pokeapi.AppState) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}
