package main

import (
	"net/http"
	"time"
	"io"
	"encoding/json"
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

func (c *Client) listLocations(pageURL *string) (LocationsResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationsResp{}, err
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResp{}, err
	}

	locations := LocationsResp{}
	err = json.Unmarshal(jsonData, &locations)
	if err != nil {
		return LocationsResp{}, err
	}

	return locations, nil
}
