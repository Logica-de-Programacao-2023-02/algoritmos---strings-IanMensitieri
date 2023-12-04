package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&entrada)

	resultado := removerEspacos(entrada)
	fmt.Printf("Resultado: %s\n", resultado)
}

func removerEspacos(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
