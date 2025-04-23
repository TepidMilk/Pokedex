package main

import (
	"time"

	"github.com/tepidmilk/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
