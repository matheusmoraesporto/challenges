package palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		word     string
		expected bool
	}{
		{
			name:     "Odd palindrome",
			word:     "renner",
			expected: true,
		},
		{
			name:     "Even palindrome",
			word:     "civic",
			expected: true,
		},
		{
			name:     "Palindrome with upper and lower case",
			word:     "Arara",
			expected: true,
		},
		{
			name:     "Not palindrome word",
			word:     "challenge",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Palindrome(tc.word)
			if tc.expected {
				require.True(t, output)
			} else {
				require.False(t, output)
			}
		})
	}
}
