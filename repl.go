package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func initREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", words[0])
	}
}
