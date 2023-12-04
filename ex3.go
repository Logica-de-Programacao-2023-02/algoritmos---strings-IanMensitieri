package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada, caractereOriginal, caractereSubstituto string

	fmt.Print("Digite uma string: ")
	fmt.Scan(&entrada)

	fmt.Print("Informe o caractere a ser substituído: ")
	fmt.Scan(&caractereOriginal)

	fmt.Print("Informe o caractere substituto: ")
	fmt.Scan(&caractereSubstituto)

	resultado := substituirCaractere(entrada, caractereOriginal, caractereSubstituto)
	fmt.Printf("Resultado: %s\n", resultado)
}

func substituirCaractere(s, original, substituto string) string {
	return strings.ReplaceAll(s, original, substituto)
}
