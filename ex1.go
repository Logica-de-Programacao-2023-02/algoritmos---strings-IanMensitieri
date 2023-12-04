package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&entrada)

	resultado := converterParaMaiusculas(entrada)
	fmt.Printf("Resultado: %s\n", resultado)
}

func converterParaMaiusculas(s string) string {
	return strings.ToUpper(s)
}
