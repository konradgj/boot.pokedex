package command

import (
	"fmt"
	"os"

	"github.com/konradgj/boot.pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	PokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show next 20 location areas",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show prev 20 location areas",
			Callback:    commandMapB,
		},
	}
}

func commandHelp(conf *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
