package repl

import (
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

func getCommands() map[string]cliCommand {
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
