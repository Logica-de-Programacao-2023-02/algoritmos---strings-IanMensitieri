package main

import (
	"fmt"
	"strings"
)

func countWords(input string) int {
	trimmedInput := strings.TrimSpace(input)

	words := strings.Fields(trimmedInput)

	return len(words)
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	numWords := countWords(userInput)

	fmt.Printf("A string contém %d palavras.\n", numWords)
}
