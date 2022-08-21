package day20

// TIL: still having some fumbles with Go's slices but getting more familiar
// For an empty slice with no capacity (for unknown growth),
//	use `var xs []int` or `xs := []int{}`
// For an "empty" slice with some capacity (for known growth):
//	use `xs := make([]int, 0, cap)` and use `xs = append(xs, x)`
//
import (
	"math"
	"valbaca.com/advent/elf"
)

const BEGIN int = 1 // optimization from Python impl, not really needed for Go

func Part1(target int) int {
	for houseNum := BEGIN; houseNum < elf.MaxInt; houseNum++ {
		facts := factors(houseNum)
		presents := elf.Sum(facts) * 10
		if presents >= target {
			return houseNum
		}
	}
	panic("No soln")
}

func Part2(target int) int {
	for houseNum := BEGIN; houseNum < elf.MaxInt; houseNum++ {
		facts := factorsWithCutoff(houseNum, houseNum/50)
		presents := elf.Sum(facts) * 11
		if presents >= target {
			return houseNum
		}
	}
	panic("No soln")
}

func factors(n int) []int {
	return factorsWithCutoff(n, 0)
}

func factorsWithCutoff(n, cutoff int) []int {
	var facts []int
	top := int(math.Ceil(math.Sqrt(float64(n))))
	for factor := 1; factor <= top; factor++ {
		mod := n % factor
		if mod == 0 {
			div := n / factor
			facts = append(facts, factor, div)
		}
	}
	if cutoff != 0 {
		filteredFacts := make([]int, 0, len(facts))
		for _, n := range facts {
			if n >= cutoff {
				filteredFacts = append(filteredFacts, n)
			}
		}
		return filteredFacts
	}
	return elf.Dedupe(facts)
}
