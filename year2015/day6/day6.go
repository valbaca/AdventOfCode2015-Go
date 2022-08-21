// Package day6
// TIL: string.FieldsFunc to split on your own custom func
// TIL: how to allocate 2D slices
package day6

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"valbaca.com/advent/elf"
)

const (
	OFF = iota
	ON
	TOGGLE
)

type Command struct {
	op int
	lx int
	rx int
	ly int
	ry int
}

func (c Command) execute(lights [][]bool) [][]bool {
	for x := c.lx; x <= c.rx; x++ {
		for y := c.ly; y <= c.ry; y++ {
			var newVal bool
			switch c.op {
			case OFF:
				newVal = false
			case ON:
				newVal = true
			case TOGGLE:
				newVal = !lights[x][y]
			}
			lights[x][y] = newVal
		}
	}
	return lights
}

func (c Command) executeDim(dims [][]uint) [][]uint {

	for x := c.lx; x <= c.rx; x++ {
		for y := c.ly; y <= c.ry; y++ {
			switch c.op {
			case OFF:
				if dims[x][y] != 0 {
					dims[x][y] -= 1
				}
			case ON:
				dims[x][y] += 1
			case TOGGLE:
				dims[x][y] += 2
			}
		}
	}
	return dims
}

func Part1(in string) string {
	lights := makeLights()
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		if s == "" {
			continue
		}
		cmd := makeCommand(s)
		lights = cmd.execute(lights)
	}
	return strconv.Itoa(countLit(lights))
}

func Part2(in string) string {
	dims := makeDims()
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		if s == "" {
			continue
		}
		cmd := makeCommand(s)
		dims = cmd.executeDim(dims)
	}
	return fmt.Sprint(countDims(dims))
}

func makeLights() [][]bool {
	lights := make([][]bool, 1000)
	for i := range lights {
		lights[i] = make([]bool, 1000)
	}
	return lights
}

func makeDims() [][]uint {
	dims := make([][]uint, 1000)
	for i := range dims {
		dims[i] = make([]uint, 1000)
	}
	return dims
}

func countDims(dims [][]uint) uint {
	var bright uint
	for _, y := range dims {
		for _, dim := range y {
			bright += dim
		}
	}
	return bright
}

func countLit(lights [][]bool) int {
	var count int
	for _, y := range lights {
		for _, light := range y {
			if light {
				count++
			}
		}
	}
	return count
}

func makeCommand(s string) Command {
	var i, op int
	if strings.HasPrefix(s, "turn off") {
		i = 9
		op = OFF
	} else if strings.HasPrefix(s, "turn on") {
		i = 8
		op = ON
	} else {
		// toggle
		i = 7
		op = TOGGLE
	}
	return makeCommandWithOp(s[i:], op)
}

func makeCommandWithOp(s string, op int) Command {
	ss := strings.FieldsFunc(s, splitPunctSpace)
	if len(ss) < 5 {
		panic("Given poor input")
	}
	lx, rx, ly, ry := elf.UnsafeAtoi(ss[0]), elf.UnsafeAtoi(ss[3]), elf.UnsafeAtoi(ss[1]), elf.UnsafeAtoi(ss[4])
	if lx > rx {
		lx, rx = rx, lx
	}
	if ly > ry {
		ly, ry = ry, ly
	}
	return Command{op, lx, rx, ly, ry}
}

func splitPunctSpace(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
}
