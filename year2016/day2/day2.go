package day2

/*
TIL: Avoiding copious bounds-checking by using graph/maps instead.
*/
import "strings"

func Part1(input string) string {
	return run(input, keypad)
}

func Part2(input string) string {
	return run(input, weirdKeypad)
}

func run(input string, keypadConfig map[string][4]string) string {
	lines := strings.Split(input, "\n")
	code := make([]string, len(lines))
	curr := "5"
	for i, line := range lines {
		for _, r := range line {
			dirIdx := directionIndexes[r]
			nxt := keypadConfig[curr][dirIdx]
			if nxt != "" {
				curr = nxt
			}
		}
		code[i] = curr
	}
	return strings.Join(code, "")
}

var (
	// UP RIGHT DOWN LEFT
	directionIndexes = map[rune]int{'U': 0, 'R': 1, 'D': 2, 'L': 3}

	keypad = map[string][4]string{
		// UP RIGHT DOWN LEFT
		//"":  {"", "", "", ""},
		"1": {"", "2", "4", ""},
		"2": {"", "3", "5", "1"},
		"3": {"", "", "6", "2"},
		"4": {"1", "5", "7", ""},
		"5": {"2", "6", "8", "4"},
		"6": {"3", "", "9", "5"},
		"7": {"4", "8", "", ""},
		"8": {"5", "9", "", "7"},
		"9": {"6", "", "", "8"},
	}

	weirdKeypad = map[string][4]string{
		"1": {"", "", "3", ""},
		"2": {"", "3", "6", ""},
		"3": {"1", "4", "7", "2"},
		"4": {"", "", "8", "3"},
		"5": {"", "6", "", ""},
		"6": {"2", "7", "A", "5"},
		"7": {"3", "8", "B", "6"},
		"8": {"4", "9", "C", "7"},
		"9": {"", "", "", "8"},
		"A": {"6", "B", "", ""},
		"B": {"7", "C", "D", "A"},
		"C": {"8", "", "", "B"},
		"D": {"B", "", "", ""},
	}
)
