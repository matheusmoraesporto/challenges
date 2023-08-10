package factorial

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		targetNum         int64
		expectedFactorial int64
	}{
		{
			targetNum:         5,
			expectedFactorial: 120,
		},
		{
			targetNum:         10,
			expectedFactorial: 3628800,
		},
		{
			targetNum:         20,
			expectedFactorial: 2432902008176640000,
		},
	}

	r := require.New(t)
	for _, tc := range testCases {
		r.Equal(tc.expectedFactorial, Factorial(tc.targetNum))
	}
}
