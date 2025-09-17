package main

import (
	"time"

	"github.com/konradgj/boot.pokedex/internal/command"
	"github.com/konradgj/boot.pokedex/internal/pokeapi"
	"github.com/konradgj/boot.pokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cfg := command.Config{
		PokeapiClient: pokeClient,
	}
	repl.Start(&cfg)
}
