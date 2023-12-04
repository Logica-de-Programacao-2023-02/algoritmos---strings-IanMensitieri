package main

import (
	"fmt"
	"unicode"
)

func isCamelCase(input string) bool {
	if len(input) == 0 || !unicode.IsUpper(rune(input[0])) {
		return false
	}

	for _, char := range input {
		if unicode.IsSpace(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return false
		}
	}

	return true
}

func countCamelCaseWords(input string) int {
	count := 1

	for _, char := range input {
		if unicode.IsUpper(char) {
			count++
		}
	}

	return count
}

func main() {
	var userInput string

	fmt.Print("Digite uma string em camelCase: ")
	fmt.Scanln(&userInput)

	if isCamelCase(userInput) {
		wordCount := countCamelCaseWords(userInput)
		fmt.Printf("A string está em camelCase e possui %d palavra(s).\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
