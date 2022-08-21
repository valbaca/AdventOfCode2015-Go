package day14

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		name      string
		expected  string
		given     string
		givenStop int
	}{
		{"", "1120", "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.", 1000},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.given, tt.givenStop)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}

}
