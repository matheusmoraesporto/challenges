package gradeaverage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGradeAverage(t *testing.T) {
	testCases := []struct {
		name        string
		grades      []float32
		expectedAvg float32
	}{
		{
			name:        "Only one grade should return grade's value",
			grades:      []float32{10.0},
			expectedAvg: 10.0,
		},
		{
			name:        "Simple test",
			grades:      []float32{5.5, 8.4, 6.1, 7.9, 10},
			expectedAvg: 7.5800004,
		},
		{
			name:        "Empty grades should return 0",
			grades:      []float32{},
			expectedAvg: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedAvg, GradeAverage(tc.grades))
		})
	}
}
