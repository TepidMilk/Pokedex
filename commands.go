package main

import "github.com/tepidmilk/pokedex/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(arg string, cfg *Config) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	current       []string
	Pokedex       map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area and display pokemon found there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catcha pokemon in current area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you've caught in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check all the pokemon you've caught!",
			callback:    commandPokedex,
		},
	}
}
