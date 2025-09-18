package repl

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.Pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
