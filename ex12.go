package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		if input[i] != input[j] {
			return false
		}
	}

	return true
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	if isPalindrome(userInput) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}
