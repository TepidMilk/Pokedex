package main

import (
	"time"

	"github.com/tepidmilk/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokeapi.NewCache(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
		pokeapiCache:  *pokeCache,
	}
	startRepl(cfg)
}
