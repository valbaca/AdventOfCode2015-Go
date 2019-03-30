package day8

import "testing"

func TestGetOverhead(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    string
	}{
		{"", 2, `""`},
		{"", 2, `"abc"`},
		{"", 3, `"aaa\"aaa"`},
		{"", 5, `"\x27"`},
		{"", 7, `"bnlxkjvl\x7docehufkj\\\"qoyhag"`},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := GetOverhead(tt.given)
			if actual != tt.expected {
				t.Errorf("`%s`: expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}

/*
  "" encodes to "\"\"", an increase from 2 characters to 6.
  "abc" encodes to "\"abc\"", an increase from 5 characters to 9.
  "\x27" encodes to "\"\\x27\"", an increase from 6 characters to 11.
*/

func TestGetAddedOverhead(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    string
	}{
		{"", 4, `""`},
		{"", 4, `"abc"`},
		{"", 5, `"\x27"`},
		{"", 6, `"aaa\"aaa"`},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := GetAddedOverhead(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}
