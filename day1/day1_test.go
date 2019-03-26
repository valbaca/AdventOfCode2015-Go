package day1

import (
	"strconv"
	"testing"
)

func TestPart1_nothing(t *testing.T) {
	a := Part1("foobar")
	if a != 0 {
		t.Error("no parens should return 0")
	}
}

func TestPart1_zero(t *testing.T) {
	a, b := Part1("()()"), Part1("(())")
	if a != 0 || b != 0 {
		t.Error("matching should be 0")
	}
}

func TestPart1_pos(t *testing.T) {
	a, b := Part1("((("), Part1("(()(()(()")
	if a != 3 || b != 3 {
		t.Error("should both be 3")
	}
}

func TestPart1_neg(t *testing.T) {
	a, b := Part1(")))"), Part1(")())())()")
	if a != -3 || b != -3 {
		t.Error("should both be -3")
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"simple", "1", ")"},
		{"simpleish", "5", "()())"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := strconv.Itoa(Part2(tt.given))
			if actual != tt.expected {
				t.Errorf("'%s': expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
