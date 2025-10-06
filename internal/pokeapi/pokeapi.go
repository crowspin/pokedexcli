package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	APIURL string = "https://pokeapi.co/api/v2"
)

func ApiGet[T any](s *AppState, url string, jsObj *T) error {
	dat, yn := s.Cache.Get(url)
	if !yn {
		//for testing cache efficacy
		//time.Sleep(2 * time.Second)

		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to send request to server")
		}
		defer res.Body.Close()

		dat, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body")
		}

		s.Cache.Add(url, dat)
	}

	if err := json.Unmarshal(dat, jsObj); err != nil {
		return fmt.Errorf("unexpected response format")
	}

	return nil
}
