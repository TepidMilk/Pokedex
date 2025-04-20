package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, value := range getCommands() {
		fmt.Sprintln("%s: %s", value.name, value.description)
	}
	return nil
}