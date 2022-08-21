package day3

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		given    string
		expected string
	}{
		{"a", ">", "2"},
		{"b", "^>v<", "4"},
		{"c", "^v^v^v^v^v", "2"},
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
func TestPart2(t *testing.T) {
	var tests = []struct {
		name     string
		given    string
		expected string
	}{
		{"a", "^v", "3"},
		{"b", "^>v<", "3"},
		{"c", "^v^v^v^v^v", "11"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Part2(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}

}
