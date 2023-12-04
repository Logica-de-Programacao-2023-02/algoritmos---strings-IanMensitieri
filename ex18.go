package main

import (
	"fmt"
	"unicode"
)

func containsOnlyNumbers(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	if containsOnlyNumbers(userInput) {
		fmt.Println("A string contém apenas números (0 a 9).")
	} else {
		fmt.Println("A string não contém apenas números (0 a 9).")
	}
}
