package main

import (
	"fmt"
	"strings"
)

func findPatternIndices(str, pattern string) []int {
	indices := []int{}
	index := -1

	for {
		index = strings.Index(str, pattern)
		if index == -1 {
			break
		}

		indices = append(indices, index)
		str = str[index+len(pattern):]
	}

	return indices
}

func main() {
	var inputString, pattern string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	fmt.Print("Digite um padrão a ser procurado na string: ")
	fmt.Scanln(&pattern)

	indices := findPatternIndices(inputString, pattern)

	if len(indices) > 0 {
		fmt.Printf("O padrão \"%s\" ocorre nos seguintes índices: %v\n", pattern, indices)
	} else {
		fmt.Printf("O padrão \"%s\" não foi encontrado na string.\n", pattern)
	}
}
