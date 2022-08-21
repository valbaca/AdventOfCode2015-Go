// Package day11
// TIL: use %c to convert byte to char in stringf
package day11

import (
	"fmt"
	"strings"
)

func Part1(in string) string {
	return findNextValidPassword(in)
}

func Part2(in string) string {
	return findNextValidPassword(Part1(in))
}

func findNextValidPassword(s string) string {
	curr := nextWord(s)
	for !validPassword(curr) {
		curr = nextWord(curr)
	}
	return curr
}

func nextWord(s string) string {
	out := make([]byte, len(s))
	copy(out, s)
	carry := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < 'a' {
			panic("invalid input")
		}
		carry = false
		if s[i] == 'z' {
			carry = true
			out[i] = 'a'
		} else {
			out[i] = s[i] + 1
		}
		if !carry {
			break
		}
	}
	if carry {
		return fmt.Sprintf("a%s", out)
	} else {
		return fmt.Sprintf("%s", out)
	}
}

func validPassword(s string) bool {
	return hasStraight(s) && hasPairs(s) && !hasAmbigious(s)
}

func hasAmbigious(s string) bool {
	return strings.ContainsAny(s, "iol")
}

func hasStraight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		now := s[i]
		nxt := s[i+1]
		lst := s[i+2]
		if now+1 == nxt && nxt+1 == lst {
			return true
		}
	}
	return false
}

func hasPairs(s string) bool {
	pairs := getPairs(s)
	return len(pairs) >= 2
}

func getPairs(s string) []string {
	pairSet := map[string]bool{}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pair := fmt.Sprintf("%c%c", s[i], s[i+1])
			pairSet[pair] = true
			i++
		}
	}
	out := []string{}
	for k := range pairSet {
		out = append(out, k)
	}
	return out
}
