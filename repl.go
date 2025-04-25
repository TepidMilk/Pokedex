package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tepidmilk/pokedex/internal/pokeapi"
)

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		commandName := text[0]
		arg := ""
		if len(text) > 1 {
			arg = text[1]
		}
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(arg, cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

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
	}
}
