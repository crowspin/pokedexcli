package pokeapi

type LocationAreaPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Link `json:"results"`
}
