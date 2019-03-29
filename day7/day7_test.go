package day7

import "testing"

func TestPart1(t *testing.T) {
	given := "123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i\n"
	var tests = []struct {
		name     string
		expected string
		givenKey string
		givenIn  string
	}{
		{"", "123", "x", given},
		{"", "456", "y", given},
		{"", "72", "d", given},    // AND
		{"", "507", "e", given},   // OR
		{"", "492", "f", given},   // LSHIFT
		{"", "114", "g", given},   // RSHIFT
		{"", "65412", "h", given}, // NOT
		{"", "65079", "i", given}, // NOT
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.givenIn, tt.givenKey)
			if actual != tt.expected {
				t.Errorf("(%q, %q): expected %s, actual %s", tt.givenIn, tt.givenKey, tt.expected, actual)
			}

		})
	}
}
