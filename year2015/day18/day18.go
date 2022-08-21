package day18

import (
	"fmt"
	"strings"
)

var cornersStuckOn bool

type Grid [][]bool

func Part1(in string, n, cycles int) string {
	grid := ParseInput(in, n)
	for i := 0; i < cycles; i++ {
		grid = grid.next()
	}
	return fmt.Sprintf("%v", grid.countOn())
}

func Part2(in string, n, cycles int) string {
	cornersStuckOn = true
	return Part1(in, n, cycles)
}

func ParseInput(in string, n int) Grid {
	grid := NewGrid(n)
	for r, line := range strings.Split(in, "\n") {
		for c, char := range line {
			grid[r][c] = (char == '#')
		}
	}
	checkStuckCorners(&grid)
	return grid
}

func NewGrid(n int) Grid {
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}
	return grid
}

func (g Grid) countOn() (count int) {
	for _, row := range g {
		for _, e := range row {
			if e {
				count++
			}
		}
	}
	return
}

func (g Grid) next() Grid {
	n := NewGrid(len(g))
	for r, row := range g {
		for c, bit := range row {
			n[r][c] = g.nextBit(bit, r, c)
		}
	}
	checkStuckCorners(&n)
	return n
}

func (g Grid) nextBit(old bool, r, c int) bool {
	neighborsOn := g.neighborsOn(r, c)
	if old && (neighborsOn == 2 || neighborsOn == 3) {
		return true
	} else if !old && neighborsOn == 3 {
		return true
	}
	return false
}

func (g Grid) neighborsOn(r, c int) (count int) {
	count += g.isNeighborOnSafe(r-1, c-1)
	count += g.isNeighborOnSafe(r-1, c)
	count += g.isNeighborOnSafe(r-1, c+1)
	count += g.isNeighborOnSafe(r, c-1)
	// self
	count += g.isNeighborOnSafe(r, c+1)
	count += g.isNeighborOnSafe(r+1, c-1)
	count += g.isNeighborOnSafe(r+1, c)
	count += g.isNeighborOnSafe(r+1, c+1)
	return
}

func (g Grid) isNeighborOnSafe(r, c int) int {
	if r >= 0 && r < len(g) && c >= 0 && c < len(g) && g[r][c] {
		return 1
	}
	return 0
}

func (g Grid) print() string {
	out := ""
	for _, row := range g {
		for _, bit := range row {
			if bit {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

func checkStuckCorners(g *Grid) {
	if cornersStuckOn {
		c := len(*g) - 1
		(*g)[0][0] = true
		(*g)[0][c] = true
		(*g)[c][0] = true
		(*g)[c][c] = true
	}
}
