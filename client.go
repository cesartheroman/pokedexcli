package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewPokeClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) listLocations(pageURL *string, cfg *config) (LocationsResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locations LocationsResp
	if entry, ok := cfg.cache.Get(url); ok {
		if err := json.Unmarshal(entry, &locations); err != nil {
			return LocationsResp{}, err
		}
		return locations, nil
	}

	fmt.Println("Cache miss for url:", url)
	res, err := http.Get(url)
	if err != nil {
		return LocationsResp{}, err
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResp{}, err
	}
	cfg.cache.Add(url, jsonData)

	err = json.Unmarshal(jsonData, &locations)
	if err != nil {
		return LocationsResp{}, err
	}

	return locations, nil
}
