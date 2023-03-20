package main

import (
	"fmt"
	"strings"
)

func bruteForce(cipherText string) {
	// Trying all possible shift values
	for i := 1; i <= 25; i++ {
		plainText := ""
		for _, ch := range cipherText {
			if ch >= 'a' && ch <= 'z' {
				plainText += string('a' + ((ch - 'a' + rune(i)) % 26))
			} else if ch >= 'A' && ch <= 'Z' {
				plainText += string('A' + ((ch - 'A' + rune(i)) % 26))
			} else {
				plainText += string(ch)
			}
		}
		fmt.Printf("Shift %d: %s\n", i, plainText)
	}
}

// Message prints in shift 3
func main() {
	cipherText := "qeb nrfzh yoltk clu grjmp lsbo qeb ixwv ald"
	//Converting Cypher to lowercase
	bruteForce(strings.ToLower(cipherText))
}
