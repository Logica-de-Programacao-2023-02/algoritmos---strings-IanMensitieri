package main

import (
	"fmt"
)

func uniqueLetters(input string) string {
	charCount := make(map[rune]int)

	for _, char := range input {
		charCount[char]++
	}

	result := ""
	for _, char := range input {
		if charCount[char] == 1 {
			result += string(char)
		}
	}

	return result
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	resultString := uniqueLetters(userInput)

	fmt.Printf("As letras únicas na string são: %s\n", resultString)
}
