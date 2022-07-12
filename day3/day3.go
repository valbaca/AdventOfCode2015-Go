// Package day3
// Santa is visiting houses and given directions with ^V<>
// Note: I first attempted this by having a House with pointers to each house,
// where the houses were initially null. But I couldn't figure out a good way to
// handle the "loop" of houses and then realized that a set of x,y coordinates
// is much easier. Sets in go are simply map[T]bool
// TIL: using map[T]bool as a set
package day3

import (
	"strconv"
)

type House struct {
	x int
	y int
}

func Part1(in string) string {
	houses := make(map[House]bool)
	curr := House{}
	houses[curr] = true
	for _, r := range in {
		if next, ok := nextHouse(curr, r); ok {
			houses[next] = true
			curr = next
		} else {
			//fmt.Printf("Got %q but ignored it\n", r) // Debug
		}
	}
	return strconv.Itoa(len(houses))
}

func Part2(in string) string {
	houses := make(map[House]bool)
	santaCurr := House{}
	roboCurr := House{}
	houses[santaCurr] = true
	var curr House
	for i, r := range in {
		if i%2 == 0 {
			curr = santaCurr
		} else {
			curr = roboCurr
		}

		if next, ok := nextHouse(curr, r); ok {
			houses[next] = true
			if i%2 == 0 {
				santaCurr = next
			} else {
				roboCurr = next
			}
		} else {
			//fmt.Printf("Got %q but ignored it\n", r) // Debug
		}

	}
	return strconv.Itoa(len(houses))
}

func nextHouse(curr House, r rune) (House, bool) {
	switch r {
	case '^':
		return House{curr.x, curr.y + 1}, true
	case '>':
		return House{curr.x + 1, curr.y}, true
	case 'v':
		return House{curr.x, curr.y - 1}, true
	case '<':
		return House{curr.x - 1, curr.y}, true
	default:
		return House{}, false
	}
}
