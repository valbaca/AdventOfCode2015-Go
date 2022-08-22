package day24

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"strings"
	"valbaca.com/advent/elf"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	ns := make([]int, len(lines))
	for i, line := range lines {
		ns[i] = elf.ParseInt(line)
	}
	pkg, ent := solve(ns, 3)
	return fmt.Sprintf("%v %v", pkg, ent)
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	ns := make([]int, len(lines))
	for i, line := range lines {
		ns[i] = elf.ParseInt(line)
	}
	pkg, ent := solve(ns, 4)
	return fmt.Sprintf("%v %v", pkg, ent)
}

func solve(ints []int, slots int) (int, int64) {
	target := elf.Sum(ints) / slots
	maxN := len(ints) - 2
	minPackages := elf.MaxInt
	var ent int64 = 0
	for i := 1; i < maxN; i++ {
		if minPackages <= i {
			return minPackages, ent
		}
		minPackages, ent = combs(ints, i, target, minPackages, ent)
	}
	return minPackages, ent
}

func combs(ints []int, i int, target int, packages int, ent int64) (int, int64) {
	viables := viableCombs(ints, i, target)
	minPackages := packages
	minEnt := ent
	for _, viable := range viables {
		if len(viable) < minPackages {
			minPackages = len(viable)
			minEnt = elf.Product(viable)
		} else if len(viable) == minPackages {
			vEnt := elf.Product(viable)
			if vEnt < minEnt {
				minPackages = len(viable)
				minEnt = vEnt
			}
		}
	}
	return minPackages, minEnt
}

func viableCombs(ints []int, i int, target int) [][]int {
	var viables [][]int
	combinations := combin.Combinations(len(ints), i)
	for _, combination := range combinations {
		potential := make([]int, i)
		for x, y := range combination {
			potential[x] = ints[y]
		}
		if elf.Sum(potential) == target {
			viables = append(viables, potential)
		}
	}
	return viables
}
