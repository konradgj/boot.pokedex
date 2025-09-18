package repl

import "fmt"

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("name required: inspect <pokemon_name>")
	}

	p, ok := cfg.Pokedex[args[0]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
