package main

import (
	"fmt"
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	reversedString := reverseString(userInput)

	fmt.Printf("A string invertida é: %s\n", reversedString)
}
