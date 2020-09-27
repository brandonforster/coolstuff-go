package stringstuff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleter(t *testing.T) {
	testCases := []struct {
		str      string
		toRemove string
		expected string
	}{
		{"zen and the art of career maintenance", "aeiou", "zn nd th rt f crr mntnnc"},
		{"zen", "puppy", "zen"},
		{"", "zen", ""},
		{"zen", "", "zen"},
	}

	for _, testCase := range testCases {
		actual := RemoveChars(testCase.str, testCase.toRemove)

		assert.Equal(t, testCase.expected, actual, "The string output should match.")
	}
}
