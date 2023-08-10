package palindrome

import (
	"fmt"
	"strings"
)

// Exercise 4: Palindrome
// Create a function that checks whether a word is a palindrome (that is, whether it reads the
// same backwards and forwards). The program must ask the user for a word and inform
// whether or not it is a palindrome.
func Palindrome(word string) bool {
	word = strings.ToLower(word)
	half := len(word) / 2
	for left := 0; left < half; left++ {
		right := len(word) - (1 + left)
		if word[left] != word[right] {
			return false
		}
	}
	return true
}

// Run is the helper function requested at exercise 4.
func Run() {
	var input string
	fmt.Println("Please provide a word to check if it is a palindrome:")
	fmt.Scan(&input)

	if Palindrome(input) {
		fmt.Printf("%q is a palindrome.\n", input)
	} else {
		fmt.Printf("%q is not a palindrome.\n", input)
	}
}
