package day6

/*
TIL: pretty easy day. Just using a Counter would've been trivial in Python.
*/
import (
	"strings"
	"unicode/utf8"

	"valbaca.com/advent/elf"
)

type Day6 struct{}

func (d Day6) Part1(input string) interface{} {
	wordLength, counts := countRunes(input)
	return runeWithHighestCount(wordLength, counts)
}

func (d Day6) Part2(input string) interface{} {
	wordLength, counts := countRunes(input)
	return runeWithLowestCount(wordLength, counts)
}

func countRunes(input string) (int, []map[rune]int) {
	lines := strings.Split(input, "\n")
	wordLength := utf8.RuneCountInString(lines[0])
	counts := make([]map[rune]int, wordLength)
	for i := 0; i < wordLength; i++ {
		counts[i] = map[rune]int{}
	}
	for _, word := range lines {
		for i, r := range word {
			counts[i][r]++
		}
	}
	return wordLength, counts
}

func runeWithHighestCount(wordLength int, counts []map[rune]int) string {
	sb := strings.Builder{}
	for i := 0; i < wordLength; i++ {
		count := counts[i]
		maxRune, maxCount := ' ', 0
		for r, n := range count {
			if n > maxCount {
				maxCount = n
				maxRune = r
			}
		}
		sb.WriteRune(maxRune)
	}
	return sb.String()
}

func runeWithLowestCount(wordLength int, counts []map[rune]int) string {
	sb := strings.Builder{}
	for i := 0; i < wordLength; i++ {
		count := counts[i]
		minRune, minCount := ' ', elf.MaxInt
		for r, n := range count {
			if n < minCount {
				minCount = n
				minRune = r
			}
		}
		sb.WriteRune(minRune)
	}
	return sb.String()
}
