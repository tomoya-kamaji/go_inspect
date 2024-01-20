package recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	// テストケースごとに期待される結果を定義します
	testCases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
		{15, 610},
	}

	// 各テストケースを実行して結果を検証します
	for _, td := range testCases {
		t.Run("Fibonacci", func(t *testing.T) {
			actual := fibonacci(td.input)
			assert.Equal(t, td.expected, actual, "fn(%d) = %d, want %d", td.input, actual, td.expected)
		})
	}
}
