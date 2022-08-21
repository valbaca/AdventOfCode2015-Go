package day11

import "testing"

func TestNextWord(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"", "b", "a"},
		{"", "aa", "z"},
		{"", "aaa", "zz"},
		{"", "ba", "az"},
		{"", "vzbxkghc", "vzbxkghb"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := nextWord(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestValidPassword(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		{"", false, "hijklmmn"},
		{"", false, "abbceffg"},
		{"", false, "abbcegjk"},
		{"", true, "abcdffaa"},
		{"", true, "ghjaabcc"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := validPassword(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}
