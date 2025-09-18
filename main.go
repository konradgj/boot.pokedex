package main

import (
	"time"

	"github.com/konradgj/boot.pokedex/internal/pokeapi"
	"github.com/konradgj/boot.pokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := repl.Config{
		PokeapiClient: pokeClient,
	}
	repl.Start(&cfg)
}
