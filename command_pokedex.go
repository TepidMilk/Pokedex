package main

import (
	"fmt"
)

func commandPokedex(arg string, cfg *Config) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.Pokedex {
		fmt.Println(" -", key)
	}
	return nil
}
