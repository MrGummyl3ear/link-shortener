package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestShorten(t *testing.T) {
	t.Run("returns an alphanumeric short identifier", func(t *testing.T) {
		type testCase struct {
			input    string
			expected string
		}

		testCases := []testCase{
			{
				input:    "google.com",
				expected: "Ln7gEz8mcI",
			},
			{
				input:    "https://www.youtube.com/watch?v=GtL1huin9EE",
				expected: "SELcdoeR9j",
			},
		}

		for _, tc := range testCases {
			actual := Hash(tc.input, 10)
			assert.Equal(t, tc.expected, actual)
		}
	})

	t.Run("is idempotent", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			assert.Equal(t, "Ln7gEz8mcI", Hash("google.com", 10))
		}
	})
}
