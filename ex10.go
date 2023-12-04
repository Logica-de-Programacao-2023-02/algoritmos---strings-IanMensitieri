package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1Chars := strings.Split(strings.ReplaceAll(str1, " ", ""), "")
	str2Chars := strings.Split(strings.ReplaceAll(str2, " ", ""), "")

	sort.Strings(str1Chars)
	sort.Strings(str2Chars)

	return strings.Join(str1Chars, "") == strings.Join(str2Chars, "")
}

func main() {
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if areAnagrams(str1, str2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
