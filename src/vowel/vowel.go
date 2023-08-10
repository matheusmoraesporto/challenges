package vowel

import (
	"fmt"
	"strings"
)

// Exercise 6: Vowel Counter
// Create a function that counts the number of vowels in a string. The program should ask the
// user for a sentence and display how many vowels it has.
func VowelCounter(str string) (count int) {
	vowels := "aeiouAEIOU"
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return
}

func Run() {
	fmt.Printf("=============================\nPlease provide a sentence to count vowels:\n\n")
	var sentence string
	fmt.Scan(&sentence)
	VowelCounter(sentence)
}
