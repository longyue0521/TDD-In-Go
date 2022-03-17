package stringcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringCalculator_Add(t *testing.T) {

	testcases := map[string]struct {
		numbers  string
		expected int
	}{
		"EmptyString_ReturnsZero": {
			numbers:  "",
			expected: 0,
		},
		"Single Number": {
			numbers:  "1",
			expected: 1,
		},
		"Another Single Number": {
			numbers:  "2",
			expected: 2,
		},
		"Two Numbers": {
			numbers:  "1,2",
			expected: 3,
		},
		"Another Two Numbers": {
			numbers:  "3,4",
			expected: 7,
		},
		"Unknown Amount Of Numbers": {
			numbers:  "1,3,5,7,9",
			expected: 25,
		},
		"Handle NewLine Delimiter": {
			numbers:  `1\n2,3`,
			expected: 6,
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			sc := NewStringCalculator()
			assert.NotNil(t, sc)
			assert.Equal(t, tt.expected, sc.Add(tt.numbers))
		})
	}

}
