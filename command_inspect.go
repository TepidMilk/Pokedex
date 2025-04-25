package main

import (
	"errors"
	"fmt"
)

func commandInspect(pokemon string, cfg *Config) error {
	if _, ok := cfg.Pokedex[pokemon]; !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", cfg.Pokedex[pokemon].Name)
	fmt.Printf("Height: %d\n", cfg.Pokedex[pokemon].Height)
	fmt.Printf("Weight: %d\n", cfg.Pokedex[pokemon].Weight)
	fmt.Println("Stats:")
	for _, stat := range cfg.Pokedex[pokemon].Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range cfg.Pokedex[pokemon].Types {
		fmt.Println(" -", t.Type.Name)
	}
	return nil
}
