package day1

import (
	"testing"
)

func TestDay1_nothing(t *testing.T) {
	a := Day1("foobar")
	if a != 0 {
		t.Error("no parens should return 0")
	}
}

func TestDay1_zero(t *testing.T) {
	a, b := Day1("()()"), Day1("(())")
	if a != 0 || b != 0 {
		t.Error("matching should be 0")
	}
}

func TestDay1_pos(t *testing.T) {
	a, b := Day1("((("), Day1("(()(()(()")
	if a != 3 || b != 3 {
		t.Error("should both be 3")
	}
}

func TestDay1_neg(t *testing.T) {
	a, b := Day1(")))"), Day1(")())())()")
	if a != -3 || b != -3 {
		t.Error("should both be -3")
	}
}
