package main

import (
	"fmt"
	"unicode"
)

func containsNumber(input string) bool {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	if containsNumber(userInput) {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}
