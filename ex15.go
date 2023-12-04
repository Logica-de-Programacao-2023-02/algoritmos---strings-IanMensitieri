package main

import (
	"fmt"
	"strings"
)

func replaceVowels(input string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			result.WriteRune('*')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	resultString := replaceVowels(userInput)

	fmt.Printf("A string com vogais substituídas é: %s\n", resultString)
}
