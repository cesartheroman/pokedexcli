package main

import (
	"time"
)

func main() {
	pokeClient := NewPokeClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

type config struct {
	pokeapiClient Client
	Next          *string
	Previous      *string
}
