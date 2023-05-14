package recursive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "階差1",
			input:    1,
			expected: 1,
		},
		{
			name:     "階差2",
			input:    2,
			expected: 2,
		},
		{
			name:     "階差5",
			input:    5,
			expected: 120,
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("Factorial: %s", td.name), func(t *testing.T) {
			actual := factorial(td.input)
			assert.Equal(t, td.expected, actual, "fn(%d) = %d, want %d", td.input, actual, td.expected)
		})
	}
}
