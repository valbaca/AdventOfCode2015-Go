package day17

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		expected string
		givenStr string
		givenInt int
	}{
		{"4", "15\n20\n10\n5\n5", 25},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			actual := Part1(tt.givenStr, tt.givenInt)
			if actual != tt.expected {
				t.Errorf("(%s, %v): expected %s, actual %s", tt.givenStr, tt.givenInt, tt.expected, actual)
			}

		})
	}

}
