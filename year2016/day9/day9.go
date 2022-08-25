package day9

/*
TIL: Now we're talking!! Faster and in fewer lines of code than the Clojure solution.
*/
import (
	"strconv"
	"strings"
	"unicode/utf8"
	"valbaca.com/advent/elf"
)

var atoi = elf.UnsafeAtoi
var atoi64 = elf.UnsafeAtoi64

func Part1(input string) string {
	count := decode(input, false)
	return strconv.FormatInt(count, 10)
}

func Part2(input string) string {
	count := decode(input, true)
	return strconv.FormatInt(count, 10)
}

func decode(s string, recur bool) (count int64) {
	length := utf8.RuneCountInString(s)
	for i := 0; i < length; i++ {
		if s[i] != '(' {
			count++
			continue
		}
		relEnd := strings.IndexRune(s[i+1:], ')')
		if relEnd == -1 {
			panic("Couldn't find ')'")
		}
		end := i + 1 + relEnd
		marker := strings.Split(s[i+1:end], "x")
		eat, times := atoi(marker[0]), atoi64(marker[1]) // eat x times

		subCount := int64(eat)
		if recur {
			eatenSubstring := s[end+1 : end+1+eat]
			subCount = decode(eatenSubstring, true)
		}

		count += subCount * times
		i = end + eat
	}
	return
}
