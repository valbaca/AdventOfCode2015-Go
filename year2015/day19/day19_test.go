package day19

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		expected   string
		givenInput string
		givenStr   string
	}{
		{"4", "H => HO\nH => OH\nO => HH", "HOH"},
		{"7", "H => HO\nH => OH\nO => HH", "HOHOHO"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			actual := Part1(tt.givenInput, tt.givenStr)
			if actual != tt.expected {
				t.Errorf("(%s, %s): expected %s, actual %s", tt.givenInput, tt.givenStr, tt.expected, actual)
			}

		})
	}

}

func TestPart2(t *testing.T) {
	transforms := "e => H\ne => O\nH => HO\nH => OH\nO => HH"
	var tests = []struct {
		expected   string
		givenInput string
		givenStr   string
	}{
		{"3", transforms, "HOH"},
		{"6", transforms, "HOHOHO"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			actual := Part2(tt.givenInput, tt.givenStr)
			if actual != tt.expected {
				t.Errorf("(%s, %s): expected %s, actual %s", tt.givenInput, tt.givenStr, tt.expected, actual)
			}

		})
	}

}
