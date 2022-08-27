package day8

import (
	"strings"
	"unicode"

	"valbaca.com/advent/elf"
)

/**
TIL: with growing familiarity, Go continues to be...very straightforward.
*/

type Day8 struct{}

func (d Day8) Part1(input string) interface{} {
	lines := elf.Lines(input)
	grid := NewGrid(50, 6)
	for _, line := range lines {
		grid = grid.ExecLine(line)
	}
	return grid.CountOn()
}

func (d Day8) Part2(input string) interface{} {
	lines := elf.Lines(input)
	grid := NewGrid(50, 6)
	for _, line := range lines {
		grid = grid.ExecLine(line)
	}
	return grid.String()
}

type Grid struct {
	grid [][]int
}

func NewGrid(wide, tall int) Grid {
	// wide is # of col, tall is # of rows
	grid := [][]int{}
	for i := 0; i < tall; i++ {
		row := make([]int, wide)
		grid = append(grid, row)
	}
	return Grid{grid}
}

func (g Grid) ExecLine(line string) Grid {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return !(unicode.IsDigit(r) || unicode.IsLetter(r))
	})
	//fmt.Printf("%v\n", fields)
	atoi := elf.UnsafeAtoi
	switch {
	case fields[0] == "rect":
		// [rect 17x1]
		ab := strings.Split(fields[1], "x")
		return g.RectOn(atoi(ab[0]), atoi(ab[1]))
	case fields[1] == "row":
		// [rotate column x 16 by 3]
		a, b := atoi(fields[3]), atoi(fields[5])
		return g.RotateRow(a, b)
	case fields[1] == "column":
		// [rotate column x 16 by 3]
		a, b := atoi(fields[3]), atoi(fields[5])
		return g.RotateCol(a, b)
	default:
		panic("Invalid line:" + line)
	}
}

func (g Grid) RectOn(a, b int) Grid {
	// a wide b tall
	for r := 0; r < b; r++ {
		for c := 0; c < a; c++ {
			g.grid[r][c] = 1
		}
	}
	return g
}

func (g Grid) RotateRow(a, b int) Grid {
	// rotate row a by b
	g.grid[a] = elf.Rotate(g.grid[a], b)
	return g
}

func (g Grid) RotateCol(a, b int) Grid {
	// rotate col a by b
	origCol := make([]int, len(g.grid))
	for i, row := range g.grid {
		origCol[i] = row[a]
	}
	rotatedCol := elf.Rotate(origCol, b)
	for i := range g.grid {
		g.grid[i][a] = rotatedCol[i]
	}
	return g
}

func (g Grid) CountOn() int {
	count := 0
	for _, row := range g.grid {
		count += elf.Sum(row)
	}
	return count
}

func (g Grid) String() string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, row := range g.grid {
		for _, i := range row {
			if i == 0 {
				sb.WriteRune('.')
			} else {
				sb.WriteRune('#')
			}
			sb.WriteRune(' ') // spacing makes easier to read
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
