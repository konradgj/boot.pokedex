package main

import "strings"

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
