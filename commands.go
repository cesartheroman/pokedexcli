package main

import (
	"errors"
	"fmt"
	"os"
)

const baseURL = "https://pokeapi.co/api/v2"

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
	locations, err := cfg.pokeapiClient.listLocations(cfg.nextLocationsURL, cfg)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("at first page")
	}

	locations, err := cfg.pokeapiClient.listLocations(cfg.prevLocationsURL, cfg)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locations.Previous
	cfg.nextLocationsURL = locations.Next

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
