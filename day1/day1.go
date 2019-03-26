package day1

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
