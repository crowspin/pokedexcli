package pokeapi

type LocationAreaData struct {
	PokemonEncounters []struct {
		Pokemon Link `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
