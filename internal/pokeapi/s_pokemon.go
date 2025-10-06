package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int  `json:"base_stat"`
		Effort   int  `json:"effort"`
		Stat     Link `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int  `json:"slot"`
		Type Link `json:"type"`
	} `json:"types"`
}
