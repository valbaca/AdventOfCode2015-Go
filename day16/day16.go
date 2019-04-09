// TIL: sometimes a map is better than copy and pasting
package day16

import (
	"fmt"
	"strings"
	"unicode"

	"valbaca.com/advent2015/utils"
)

var outdatedRetroencabulator bool

type Aunt struct {
	number      int
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func Part1(in string) string {
	outdatedRetroencabulator = false
	return FindMatchingAunt(in)
}
func Part2(in string) string {
	outdatedRetroencabulator = true
	return FindMatchingAunt(in)
}
func FindMatchingAunt(in string) string {
	aunts := ParseInput(in)
	ticker := Aunt{-1, 3, 7, 2, 3, 0, 0, 5, 3, 2, 1}
	for _, aunt := range aunts {
		if ticker.matches(aunt) {
			return fmt.Sprintf("%v", aunt.number)
		}
	}
	return "no aunt found"
}

func ParseInput(in string) []Aunt {
	aunts := []Aunt{}
	for _, line := range strings.Split(in, "\n") {
		aunts = append(aunts, ParseLine(line))
	}
	return aunts
}

func ParseLine(line string) Aunt {
	// Sue 1: goldfish: 6, trees: 9, akitas: 0
	// 0   1  2         3  4      5  6       7
	aunt := Aunt{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	split := strings.FieldsFunc(line, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})
	aunt.number = utils.Atoi(split[1])
	for i := 2; i < len(split)-1; i += 2 {
		key, val := split[i], utils.Atoi(split[i+1])
		if key == "children" {
			aunt.children = val
		}
		if key == "cats" {
			aunt.cats = val
		}
		if key == "samoyeds" {
			aunt.samoyeds = val
		}
		if key == "pomeranians" {
			aunt.pomeranians = val
		}
		if key == "akitas" {
			aunt.akitas = val
		}
		if key == "vizslas" {
			aunt.vizslas = val
		}
		if key == "goldfish" {
			aunt.goldfish = val
		}
		if key == "trees" {
			aunt.trees = val
		}
		if key == "cars" {
			aunt.cars = val
		}
		if key == "perfumes" {
			aunt.perfumes = val
		}
	}
	return aunt
}

func (a Aunt) matches(other Aunt) bool {
	if other.children != -1 && other.children != a.children {
		return false
	}
	if outdatedRetroencabulator {
		if other.cats != -1 && other.cats <= a.cats {
			return false
		}
	} else if other.cats != -1 && other.cats != a.cats {
		return false
	}

	if other.samoyeds != -1 && other.samoyeds != a.samoyeds {
		return false
	}
	if outdatedRetroencabulator {
		if other.pomeranians != -1 && other.pomeranians >= a.pomeranians {
			return false
		}
	} else if other.pomeranians != -1 && other.pomeranians != a.pomeranians {
		return false
	}
	if other.akitas != -1 && other.akitas != a.akitas {
		return false
	}
	if other.vizslas != -1 && other.vizslas != a.vizslas {
		return false
	}
	if outdatedRetroencabulator {
		if other.goldfish != -1 && other.goldfish >= a.goldfish {
			return false
		}
	} else if other.goldfish != -1 && other.goldfish != a.goldfish {
		return false
	}
	if outdatedRetroencabulator {
		if other.trees != -1 && other.trees <= a.trees {
			return false
		}
	} else if other.trees != -1 && other.trees != a.trees {
		return false
	}
	if other.cars != -1 && other.cars != a.cars {
		return false
	}
	if other.perfumes != -1 && other.perfumes != a.perfumes {
		return false
	}
	return true
}
