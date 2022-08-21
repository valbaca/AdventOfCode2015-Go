package day12

import "testing"

func TestSumString(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"", "42", "42"},
		{"", "42", "[[[42]]]"},
		{"", "42", "{{42}}"},
		{"", "-42", "-42"},
		{"", "-41", "1,-42"},
		/*
			[1,2,3] and {"a":2,"b":4} both have a sum of 6.
			[[[3]]] and {"a":{"b":4},"c":-1} both have a sum of 3.
			{"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0.
			[] and {} both have a sum of 0.
		*/
		{"", "6", `[1,2,3]`},
		{"", "6", `{"a":2,"b":4}`},
		{"", "3", `[[[3]]]`},
		{"", "3", `{"a":{"b":4},"c":-1}`},
		{"", "0", `{"a":[-1,1]}`},
		{"", "0", `[-1,{"a":1}]`},
		{"", "0", `[]`},
		{"", "0", `{}`},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := sumString(tt.given)
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
		{"", "6", "[1,2,3]"},
		{"", "4", `[1,{"c":"red","b":2},3]`},
		{"", "0", `{"d":"red","e":[1,2,3,4],"f":5}`},
		{"", "6", `[1,"red",5]`},
		{"", "6", `{"a":[1,"red",5]}`},
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
