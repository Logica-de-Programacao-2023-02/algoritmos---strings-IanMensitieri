package main

import (
	"fmt"
	"strings"
)

func removeVowels(input string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range input {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	resultString := removeVowels(userInput)

	fmt.Printf("A string sem vogais é: %s\n", resultString)
}
