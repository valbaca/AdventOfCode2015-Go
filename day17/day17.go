// TIL: While I initially thought constructing all the permutations would be
// slow and difficult, once I realized I could use a bit set it got easier.
// Then just deciding to go through 0..2**n and making sets from that Finally,
// I finished by realizing I didn't need the sets, just their sum (for part 1)
// and their length (for part 2)
package day17

import (
	"fmt"
	"strings"

	"valbaca.com/advent2015/utils"
)

func Part1(in string, total int) string {
	buckets := ParseInput(in)
	return fmt.Sprintf("%v", FindNumCombos(buckets, total))
}

func Part2(in string, total int) string {
	buckets := ParseInput(in)
	return fmt.Sprintf("%v", FindNumMinCombos(buckets, total))
}

func ParseInput(in string) []int {
	out := []int{}
	for _, line := range strings.Split(in, "\n") {
		out = append(out, utils.Atoi(line))
	}
	return out
}

func FindNumCombos(buckets []int, total int) int {
	count := 0
	for i := 0; i < (1 << uint(len(buckets))); i++ {
		sum, _ := SumSubset(buckets, i)
		if sum == total {
			count++
		}
	}
	return count
}

func FindNumMinCombos(buckets []int, total int) int {
	count := 0
	minNumBucket := utils.MaxInt
	for i := 0; i < (1 << uint(len(buckets))); i++ {
		sum, setLen := SumSubset(buckets, i)
		if sum == total {
			if setLen < minNumBucket {
				// new lowest count
				minNumBucket = setLen
				count = 1
			} else if setLen == minNumBucket {
				count++
			}
		}
	}
	return count
}

// This was my original approach, but I realized I didn't need the actual subsets
// and appending arrays was expensive (or would be IRL)
// So I used SumSubset instead, which sped up runtime from 0.67s to 0.19s
func MakeSubSet(buckets []int, bits int) []int {
	out := []int{}
	for _, x := range buckets {
		bit := bits % 2
		bits >>= 1
		if bit == 1 {
			out = append(out, x)
		}
	}
	return out
}

// Faster than MakeSubSet by just doing the math inline and not creating a slice
func SumSubset(buckets []int, bits int) (sum, length int) {
	for _, x := range buckets {
		if bits%2 == 1 {
			sum += x
			length++
		}
		bits >>= 1
	}
	return
}
