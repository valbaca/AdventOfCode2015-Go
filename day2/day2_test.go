package day2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"ex1", "58", "2x3x4\n"},
		{"ex2", "43", "1x1x10\n"},
		{"ex3", "101", "2x3x4\n1x1x10\n"},
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
		expected string
		given    string
	}{
		{"ex1", "34", "2x3x4\n"},
		{"ex2", "14", "1x1x10\n"},
		{"ex3", "48", "2x3x4\n1x1x10\n"},
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
