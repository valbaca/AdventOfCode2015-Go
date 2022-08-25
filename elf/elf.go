// Package elf: elves are Santa's little helpers!
// Utility functions, in particular short, unsafe versions of functions useful for advent solns
package elf

import (
	"bufio"
	"strconv"
	"strings"
)

// c/o https://stackoverflow.com/a/6878625
// I've flipped the order; it shows how they're derived

const MinUint = 0        // 000...
const MaxUint = ^uint(0) // 111....

const MaxInt = int(MaxUint >> 1) // 0111....
const MinInt = -MaxInt - 1       // 1000..

// UnsafeAtoi is strconv.Atoi that simply panics on any error
func UnsafeAtoi(s string) int {
	if out, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return out
	}
}

func UnsafeAtoi64(s string) int64 {
	if out, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	} else {
		return out
	}
}

func UnsafeAtoUint16(s string) uint16 {
	if out, err := strconv.ParseUint(s, 10, 16); err != nil {
		panic(err)
	} else {
		return uint16(out)
	}
}

func Max(ints ...int) int {
	max := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}

func ParseInt(s string) int {
	if strings.HasSuffix(s, ",") {
		s = s[:len(s)-1]
	}
	return UnsafeAtoi(s)
}

func Sum(a []int) int {
	s := 0
	for _, n := range a {
		s += n
	}
	return s
}

func Product(a []int) int64 {
	var p int64 = 1
	for _, n := range a {
		p *= int64(n)
	}
	return p
}

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func Dedupe(a []int) []int {
	set := make(map[int]bool)
	for _, n := range a {
		set[n] = true
	}
	// in The Future, may be able to just: maps.Keys(set) but avoiding /x/exp imports...for now.
	deduped := make([]int, len(set))
	i := 0
	for key := range set {
		deduped[i] = key
		i++
	}
	return deduped
}

func SplitWords(s string) []string {
	// https://pkg.go.dev/bufio#example-Scanner-Words
	// https://stackoverflow.com/questions/38026530/a-better-way-to-use-scanner-for-multiple-tokens-per-line
	var splits []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		splits = append(splits, word)
	}
	return splits
}

func Lines(s string) []string {
	return strings.Split(s, "\n")
}

func CountLines(lines []string, test func(s string) bool) int {
	count := 0
	for _, line := range lines {
		if test(line) {
			count++
		}
	}
	return count
}

// Rotate array to the right by n. To rotate to the left, use negative n.
func Rotate(a []int, n int) []int {
	out := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		j := Mod(i+n, len(a))
		out[j] = a[i]
	}
	return out
}

// Mod performs the modulus operation, x mod y.
// The answer is always within the range of 0...|y| and has the sign of y
// Use this when you want a safe value within a range...like with indexing, i.e. a[-1] == a[len(a)-1]
// Note: many people pronounce `%` as "mod" but it's actually "remainder" or "rem"
// See https://clojuredocs.org/clojure.core/mod
func Mod(x, y int) int {
	if x == 0 {
		return 0
	}
	rem := x % y
	if Sign(x) == Sign(y) {
		return rem
	} else {
		return rem + y
	}
}

// Sign returns 1, 0, or -1 if n is positive, zero, or negative
func Sign(n int) int {
	if n > 0 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return -1
	}
}
