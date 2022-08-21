package day18

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		expected    string
		givenStart  string
		givenN      int
		givenCycles int
	}{
		{"4", ".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####..", 6, 4},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			actual := Part1(tt.givenStart, tt.givenN, tt.givenCycles)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.givenStart, tt.expected, actual)
			}

		})
	}

}
