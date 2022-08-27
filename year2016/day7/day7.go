package day7

/*
TIL: the string/ranges/runes/int still gets in the way.
While Go still makes it challenging to introduce higher-order functions due to its limited type system, I am starting to
get used to it.
It does have an interesting ironic side-benefit of ensuring that I do NOT spend more time adding unnecessary abstractions.
This was something I ran afoul a lot in Clojure. Because types were dynamic and functions were so easily composed, it
gave direct power, but also led me to feel the need to repeatedly refactor and abstract every day's solution until it
was just a few lines.
*/
import (
	"unicode/utf8"

	"valbaca.com/advent/elf"
)

type Day7 struct{}

func (d Day7) Part1(input string) interface{} {
	return elf.CountLines(elf.Lines(input), supportsTLS)

}

func (d Day7) Part2(input string) interface{} {
	return elf.CountLines(elf.Lines(input), supportsSSL)

}

func supportsTLS(s string) bool {
	// cannot use strings.FieldsFunc because it requires the f be stateless
	end := utf8.RuneCountInString(s) - 3
	withinBrackets := false
	hasAbbaOutside := false
	hasAbbaInside := false
	for i, r := range s {
		if i >= end {
			break
		}
		if r == '[' {
			withinBrackets = true
			continue
		}
		if r == ']' {
			withinBrackets = false
			continue
		}
		if r == rune(s[i+3]) && r != rune(s[i+1]) && s[i+1] == s[i+2] {
			if withinBrackets {
				hasAbbaInside = true
			} else {
				hasAbbaOutside = true
			}
		}
	}
	return hasAbbaOutside && !hasAbbaInside
}

func supportsSSL(s string) bool {
	end := utf8.RuneCountInString(s) - 2
	withinBrackets := false
	abaOutsideWanted := []string{}
	babInsideWanted := []string{}
	for i, r := range s {
		if i >= end {
			break
		}
		if r == '[' {
			withinBrackets = true
			continue
		}
		if r == ']' {
			withinBrackets = false
			continue
		}
		if r == rune(s[i+2]) && r != rune(s[i+1]) {
			if withinBrackets {
				// we have bab, so see if it's wanted, otherwise add aba to wanted
				bab := s[i : i+3]
				for _, babWanted := range babInsideWanted {
					if bab == babWanted {
						return true
					}
				}
				aba := string([]rune{rune(s[i+1]), rune(s[i]), rune(s[i+1])})

				abaOutsideWanted = append(abaOutsideWanted, aba)
			} else {
				// here we do the literal opposite
				aba := s[i : i+3]
				for _, abaWanted := range abaOutsideWanted {
					if aba == abaWanted {
						return true
					}
				}
				bab := string([]rune{rune(s[i+1]), rune(s[i]), rune(s[i+1])})
				babInsideWanted = append(babInsideWanted, bab)
			}
		}
	}
	return false
}
