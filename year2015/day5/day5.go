// Package day5
// TIL: use ... for inline array lengths
// TIL: can't define constant maps or arrays. must use var
// TIL: build your own regexs with strings
package day5

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(in string) string {
	count := 0
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		if isNice(s) {
			count++
		}
	}
	return strconv.Itoa(count)
}

var vowels = [...]rune{'a', 'e', 'i', 'o', 'u'}

var naughtyPrev = map[rune]rune{
	// curr : prev
	'b': 'a',
	'd': 'c',
	'q': 'p',
	'y': 'x',
}

func isNice(s string) bool {
	if len(s) < 3 {
		return false
	}
	vowels := 0
	dub := false
	var p rune
	for i, r := range s {
		if vowels < 3 && isVowel(r) {
			vowels++
		}
		if i == 0 {
			p = r
			continue
		}
		if !dub && p == r {
			dub = true
		}
		if isNaughtyPair(p, r) {
			return false
		}
		p = r
	}
	return vowels >= 3 && dub
}

func isVowel(r rune) bool {
	for _, v := range vowels {
		if v == r {
			return true
		}
	}
	return false
}

func isNaughtyPair(p, r rune) bool {
	return p == naughtyPrev[r]
}

func Part2(in string) string {
	count := 0
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		if isNicer(s) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func isNicer(s string) bool {
	return hasPair(s) && hasDouble(s)
}

func hasPair(s string) bool {
	if len(s) < 4 {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		j := i + 1
		a := string(s[i])
		b := string(s[j])
		re, _ := regexp.Compile("" + a + b + ".*" + a + b)
		if re.MatchString(s) {
			return true
		}
	}
	return false
}

func hasDouble(s string) bool {
	if len(s) < 3 {
		return false
	}
	for i := 0; i < len(s)-2; i++ {
		a := string(s[i])
		re, _ := regexp.Compile("" + a + "." + a)
		if re.MatchString(s) {
			return true
		}
	}
	return false
}
