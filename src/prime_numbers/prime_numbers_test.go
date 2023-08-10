package primenumbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimeNumber(t *testing.T) {
	testCases := []struct {
		name           string
		targetNumber   int
		expectedResult bool
	}{
		{
			name:           "Simple prime number",
			targetNumber:   5,
			expectedResult: true,
		},
		{
			name:           "High prime number",
			targetNumber:   997,
			expectedResult: true,
		},
		{
			name:           "Not prime number",
			targetNumber:   10,
			expectedResult: false,
		},
		{
			name:           "Not primer number case lower than 1",
			targetNumber:   0,
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedResult, PrimeNumbers(tc.targetNumber))
		})
	}
}
