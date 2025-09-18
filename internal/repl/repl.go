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

		parts := cleanInput(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}

		err := cmd.Callback(cfg, args...)
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
