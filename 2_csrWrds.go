package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Execution(wordInputStr string, notAllowed []string) string {
	var newInputStr []string

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	newSlice := strings.FieldsFunc(wordInputStr, f)
	newSlice1 := strings.Fields(wordInputStr)

	for pos1, wordInputStr := range newSlice {
		for _, notAllowedWord := range notAllowed {
			if wordInputStr == notAllowedWord {
				insert := strings.Repeat("*", len(notAllowedWord))
				newSlice1[pos1] = insert
				newInputStr = append(newSlice1[:pos1+1])
			}
		}
	}
	return strings.Join(newInputStr, " ")
}

func main() {

	var notAllowed = []string{"dolor", "elit", "quis", "nisi", "fugiat", "proident", "laborum"}
	inputStr := "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	fmt.Println("[Original : ]", inputStr)
	E := Execution(inputStr, notAllowed)
	fmt.Println("[Censored : ]", E)
}
