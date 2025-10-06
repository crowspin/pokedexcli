package main

import (
	"fmt"

	"github.com/crowspin/pokeapi"
)

func commandMap(c *pokeapi.AppState) {
	fetchMapPage(c, c.Paginate.Next)
}

func commandMapb(c *pokeapi.AppState) {
	fetchMapPage(c, c.Paginate.Previous)
}

func fetchMapPage(c *pokeapi.AppState, dest string) {
	if dest == "" {
		dest = pokeapi.APIURL + "/location-area"
	}

	var js pokeapi.LocationAreaPage
	if err := pokeapi.ApiGet(c, dest, &js); err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.Paginate.Next = js.Next
	c.Paginate.Previous = js.Previous

	for _, val := range js.Results {
		fmt.Println(val.Name)
	}
}
