package main

import (
	"fmt"
	"strings"
)

func replaceLetter(input string, oldChar, newChar string) string {
	return strings.ReplaceAll(input, oldChar, newChar)
}

func main() {
	var userInput, oldChar, newChar string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&newChar)

	resultString := replaceLetter(userInput, oldChar, newChar)

	fmt.Printf("A string após a substituição é: %s\n", resultString)
}
