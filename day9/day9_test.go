package day9

import "testing"

func TestSplitLine(t *testing.T) {
	var tests = []struct {
		name     string
		expected Edge
		given    string
	}{
		{"", Edge{"London", "Dublin", 464}, "London to Dublin = 464"},
		{"", Edge{"London", "Dublin", 464}, "London to Dublin = 464\n"},
		{"", Edge{"London", "Belfast", 518}, "London to Belfast = 518"},
		{"", Edge{"Dublin", "Belfast", 141}, "Dublin to Belfast = 141"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := SplitLine(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %+v, actual %+v", tt.given, tt.expected, actual)
			}

		})
	}
}

/*
func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"", "605", "London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141\n"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
*/
