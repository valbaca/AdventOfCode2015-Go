package day10

import "testing"

func TestLookAndSay(t *testing.T) {
	var tests = []struct {
		name     string
		given    string
		expected string
	}{
		{"", "1", "11"},          // one "one"
		{"", "11", "21"},         // two "one"
		{"", "21", "1211"},       // one "two", one one
		{"", "1211", "111221"},   // one "one", one "two", two "one"
		{"", "111221", "312211"}, // three "one", two "two", one "one"
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := lookAndSay(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}

}
