package table

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTable(t *testing.T) {
	expectedTable := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	receivedTable := Table(10)
	require.Equal(t, expectedTable, receivedTable)
}
