package simplecalculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleCalculator(t *testing.T) {
	testCases := []struct {
		name             string
		op               string
		x                int
		y                int
		expectedResult   int
		expectedErrorMsg string
	}{
		{
			name:           "Sum operation",
			op:             "+",
			x:              25,
			y:              30,
			expectedResult: 55,
		},
		{
			name:           "Negative sum operation",
			op:             "+",
			x:              -25,
			y:              30,
			expectedResult: 5,
		},
		{
			name:           "Subtraction operation",
			op:             "-",
			x:              100,
			y:              5,
			expectedResult: 95,
		},
		{
			name:           "Negative subtraction operation",
			op:             "-",
			x:              100,
			y:              -5,
			expectedResult: 105,
		},
		{
			name:           "Multiplication operation",
			op:             "*",
			x:              100,
			y:              5,
			expectedResult: 500,
		},
		{
			name:           "Division operation",
			op:             "/",
			x:              100,
			y:              5,
			expectedResult: 20,
		},
		{
			name:             "Division operation should return error when divisor is 0",
			op:               "/",
			x:                5,
			y:                0,
			expectedErrorMsg: "the divisor can not be equal to 0",
		},
		{
			name:             "Should return error when the operator is invalid",
			op:               "%",
			x:                1,
			y:                1,
			expectedErrorMsg: "the operator \"%\" is not valid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)
			receivedResult, receivedErr := SimpleCalculator(tc.x, tc.y, tc.op)
			if tc.expectedErrorMsg != "" {
				r.EqualError(receivedErr, tc.expectedErrorMsg)
			} else {
				r.Equal(tc.expectedResult, receivedResult)
			}
		})
	}
}
