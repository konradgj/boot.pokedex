package repl

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("name required: catch <pokemon_name>")
	}

	resp, err := cfg.PokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	caught := attempCatch(resp.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	if caught {
		fmt.Printf("%s was caught!\n", resp.Name)
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}

	cfg.Pokedex[resp.Name] = resp
	return nil
}

func attempCatch(baseExp int) bool {
	chance := normalizeBaseExp(baseExp)
	if chance > 95 {
		chance = 95
	}
	return rand.Float64() < chance
}

func normalizeBaseExp(baseExp int) float64 {
	minExp := 20.0
	maxExp := 255.0

	n := (maxExp - float64(baseExp)) / (maxExp - minExp)

	if n < 0 {
		return 0
	}
	if n > 1 {
		return 1
	}
	return n
}
