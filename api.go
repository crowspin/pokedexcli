package main

type Link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Link `json:"results"`
}

type LocationAreaData struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod Link `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int  `json:"rate"`
			Version Link `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location Link `json:"location"`
	Names    []struct {
		Name     string `json:"name"`
		Language Link   `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        Link `json:"pokemon"`
		VersionDetails []struct {
			Version          Link `json:"version"`
			MaxChance        int  `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int  `json:"min_level"`
				MaxLevel int  `json:"max_level"`
				Chance   int  `json:"chance"`
				Method   Link `json:"method"`
				//ConditionValues []struct{?}
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
