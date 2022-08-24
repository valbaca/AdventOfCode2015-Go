package day1

import (
	"strconv"
	"valbaca.com/advent/elf"
)

/*
TIL: cannot use a slice as a map key, so had to create a `pos` struct instead. Considered using a string.
Had to write my own Abs function...getting deja-vu that I've done this before...
*/
import (
	"strings"
	"unicode"
)

func Part1(input string) string {
	dirs := parse(input)
	x, y, facing := 0, 0, N
	turn := map[string]int{"L": -1, "R": 1}
	for _, dir := range dirs {
		facing = rotate(facing, turn[dir.lr])
		x, y = move(x, y, facing, dir.steps)
	}
	manhDist := elf.Abs(x) + elf.Abs(y)
	return strconv.Itoa(manhDist)
}

func Part2(input string) string {
	dirs := parse(input)
	x, y, facing := 0, 0, N
	turn := map[string]int{"L": -1, "R": 1}
	seen := map[pos]bool{}
	seen[pos{x, y}] = true
	for _, dir := range dirs {
		facing = rotate(facing, turn[dir.lr])
		for i := 0; i < dir.steps; i++ {
			x, y = move(x, y, facing, 1)
			ps := pos{x, y}
			if _, hasSeen := seen[ps]; hasSeen {
				manhDist := elf.Abs(x) + elf.Abs(y)
				return strconv.Itoa(manhDist)
			} else {
				seen[ps] = true
			}
		}

	}
	manhDist := elf.Abs(x) + elf.Abs(y)
	return strconv.Itoa(manhDist)
}

type inst struct {
	lr    string
	steps int
}

func parse(input string) []inst {
	xs := strings.FieldsFunc(input, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})
	dirs := make([]inst, len(xs))
	for i := 0; i < len(xs); i++ {
		dirs[i] = inst{string(xs[i][0]), elf.UnsafeAtoi(xs[i][1:])}
	}
	return dirs
}

const (
	N = iota
	E
	S
	W
)

func rotate(direction int, times int) int {
	direction += times
	if direction < 0 {
		direction %= 4
		direction += 4
	} else if direction > W {
		direction %= 4
	}
	return direction
}

func move(x, y, card, distance int) (int, int) {
	switch card {
	case N:
		return x, y - distance
	case E:
		return x + distance, y
	case S:
		return x, y + distance
	case W:
		return x - distance, y
	default:
		panic("Invalid direction")
	}
}

type pos struct {
	x, y int
}
