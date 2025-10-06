package pokeapi

import "github.com/crowspin/pokecache"

type AppState struct {
	Paginate struct {
		Next     string
		Previous string
	}
	Arg     string
	Pokedex map[string]Pokemon
	Cache   pokecache.Cache
}
