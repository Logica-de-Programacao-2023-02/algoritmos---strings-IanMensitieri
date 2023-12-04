package main

import (
	"fmt"
	"strconv"
)

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	var userInput string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&userInput)

	if isFloat(userInput) {
		fmt.Printf("%s é um número válido em ponto flutuante.\n", userInput)
	} else {
		fmt.Printf("%s não é um número válido em ponto flutuante.\n", userInput)
	}
}
