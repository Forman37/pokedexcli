package main

import(
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)

	returnStrings := []string{}
	for _, word := range words {
		returnStrings = append(returnStrings, strings.ToLower(strings.TrimSpace(word)))
	}
	return returnStrings
}
