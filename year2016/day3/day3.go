package day3

/*
TIL: the "partition" part of this took way longer than expected. Just that function took more lines than the entire
Clojure solution took...
On the other hand, running the code takes *far* less time than it takes for the Clojure REPL to start
*/
import (
	"sort"
	"strconv"
	"strings"
	"valbaca.com/advent/elf"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		sides := make([]int, 3)
		for i := 0; i < 3; i++ {
			sides[i] = elf.UnsafeAtoi(fields[i])
		}
		if isTriangle(sides) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	count := 0
	allSides := partitionSides(lines)
	for _, sides := range allSides {
		if isTriangle(sides) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func isTriangle(sides []int) bool {
	sort.Ints(sides)
	return sides[2] < (sides[0] + sides[1])
}

func partitionSides(lines []string) (allSides [][]int) {
	// deal with lines one 3x3 grid at a time...
	for l := 0; l < len(lines); l += 3 {
		var trioInts []int
		for _, trioLine := range lines[l : l+3] {
			fields := strings.Fields(trioLine)
			for _, field := range fields {
				trioInts = append(trioInts, elf.UnsafeAtoi(field))
			}
		}
		for i := 0; i < 3; i++ {
			var sides []int
			for j := 0; j < 3; j++ {
				// this is where the "flip" happens, from row-based to col-based
				// normally you would do [i*rowLen + j], here we go "down" the column faster with j
				sides = append(sides, trioInts[j*3+i])
			}
			allSides = append(allSides, sides)
		}
	}
	return allSides
}
