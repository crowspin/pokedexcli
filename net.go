package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/crowspin/pokecache"
)

const (
	APIURL string = "https://pokeapi.co/api/v2"
)

func apiGet(endpoint string) ([]byte, error) {
	var v pokecache.Cache
	url := APIURL + endpoint

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to send request to server")
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response body")
	}

	return dat, nil
}
