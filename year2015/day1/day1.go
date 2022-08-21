// Package day1
// Santa goes up and down an elevator with '(' for ups, and ')' for downs
// TIL: runes are *like* chars in other langs. Single quotes for runes
package day1

// Part1 gives the floor that santa lands on based on open/closed parenthesis
func Part1(in string) int {
	out := 0
	for _, r := range in {
		if r == '(' {
			out++
		} else if r == ')' {
			out--
		}
	}
	return out
}

// Part2 (per usual) is almost just like Part1, but we can exit early and just
// need to return the index
func Part2(in string) int {
	out := 0
	for i, r := range in {
		if r == '(' {
			out++
		} else if r == ')' {
			out--
		}
		if out == -1 {
			return i + 1
		}
	}
	panic("never went to floor -1")
}
