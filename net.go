package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	APIURL string = "https://pokeapi.co/api/v2"
)

func apiGet(url string) ([]byte, error) {
	if val, yn := cache.Get(url); yn {
		return val, nil
	}

	//time.Sleep(2 * time.Second)

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to send request to server")
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response body")
	}

	cache.Add(url, dat)

	return dat, nil
}
