package repl

import "fmt"

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("name required: explore <area_name>")
	}

	resp, err := cfg.PokeapiClient.GetLocatioArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, encounters := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}

	return nil
}
