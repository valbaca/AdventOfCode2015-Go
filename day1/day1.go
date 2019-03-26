package day1

func Day1(in string) int {
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
