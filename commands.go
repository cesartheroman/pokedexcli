package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/"

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

func commandMapf(cfg *config) error {
	url := baseURL
	if cfg.Next != nil {
		url += fmt.Sprintf("?%s", *cfg.Next)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations LocationsResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return err
	}

	var nextOffset string
	var prevOffset string
	if next := locations.Next; next != nil {
		nextOffset = strings.Split(*next, "?")[1]
	}
	if prev := locations.Previous; prev != nil {
		prevOffset = strings.Split(*prev, "?")[1]
	}

	cfg.Next = &nextOffset
	cfg.Previous = &prevOffset

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	url := baseURL
	if cfg.Previous == nil || *cfg.Previous == "" {
		return errors.New("at first page")
	}
	url += fmt.Sprintf("/?%s", *cfg.Previous)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations LocationsResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return err
	}

	var nextOffset string
	var prevOffset string
	if next := locations.Next; next != nil {
		nextOffset = strings.Split(*next, "?")[1]
	}
	if prev := locations.Previous; prev != nil {
		prevOffset = strings.Split(*prev, "?")[1]
	}

	cfg.Previous = &prevOffset
	cfg.Next = &nextOffset

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

type LocationsResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
