package main

import (
	"time"
)

func main() {
	pokeClient := NewPokeClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache: NewCache(30 * time.Second),
	}
	startRepl(cfg)
}

type config struct {
	pokeapiClient    Client
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *Cache
}
