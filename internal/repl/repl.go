package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, exists := getCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}

		err := cmd.Callback(cfg)
		if err != nil && err.Error() == "exit" {
			fmt.Println("Closing the Pokedex... Goodbye!")
			break
		} else if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return nil
	}
	parts := strings.Fields(text)
	for i, word := range parts {
		parts[i] = strings.ToLower(word)
	}
	return parts
}
