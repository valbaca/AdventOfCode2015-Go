package day4

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"ex1", "609043", "abcdef"},
		{"ex2", "1048970", "pqrstuv"},
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

func TestGetHash(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		givenStr string
		givenInt int
	}{
		{"ex1", "000001dbbfa3a5c83a2d506429c7b00e", "abcdef", 609043},
		{"ex2", "000006136ef2ff3b291c85725f17325c", "pqrstuv", 1048970},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := GetHash(tt.givenStr, tt.givenInt)
			if actual != tt.expected {
				t.Errorf("GetHash(%s%d): expected %s, actual %s", tt.givenStr, tt.givenInt, tt.expected, actual)
			}

		})
	}
}
