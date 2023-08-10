package vowel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVowelCounter(t *testing.T) {
	testCases := []struct {
		name           string
		targetWord     string
		expectedOutput int
	}{
		{
			name:           "Simple test",
			targetWord:     "photography",
			expectedOutput: 3,
		},
		{
			name:           "Word only with vowels",
			targetWord:     "aeiouAEIOU",
			expectedOutput: 10,
		},
		{
			name:           "Word without vowel",
			targetWord:     "Dry",
			expectedOutput: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedOutput, VowelCounter(tc.targetWord))
		})
	}
}
