package main

import (
	"fmt"
	"strconv"
)

func isIncreasingSequence(input string) bool {
	if len(input) < 2 {
		return false
	}

	for i := 1; i < len(input); i++ {
		prevDigit, _ := strconv.Atoi(string(input[i-1]))
		currentDigit, _ := strconv.Atoi(string(input[i]))

		if prevDigit >= currentDigit {
			return false
		}
	}

	return true
}

func main() {
	var userInput string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&userInput)

	if isIncreasingSequence(userInput) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
