package cache_test

import (
	"link-shortener/internal/storage/cache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveGet(t *testing.T) {
	p := cache.NewInMemory()
	t.Run("returns a shorten link", func(t *testing.T) {
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
			err := p.Save(tc.input, tc.expected)
			if err != nil {
				t.Errorf("error occured: %v", err)
			}
			actual, err := p.Get(tc.expected)
			if err != nil {
				t.Errorf("error occured: %v", err)
			}
			assert.Equal(t, tc.input, actual)
		}
	})
}
