// utils provides short, unsafe versions of functions for use in hacking
package utils

import (
	"strconv"
	"strings"
)

// c/o https://stackoverflow.com/a/6878625
// I've flipped the order so it shows how they're derived
const MinUint = 0        // 000...
const MaxUint = ^uint(0) // 111....

const MaxInt = int(MaxUint >> 1) // 0111....
const MinInt = -MaxInt - 1       // 1000..

// AtoI is an unsafe verison of strconv.Atoi
// it panics if there's any error
func AtoI(s string) int {
	if out, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return out
	}
}

func AtoUint16(s string) uint16 {
	if out, err := strconv.ParseUint(s, 10, 16); err != nil {
		panic(err)
	} else {
		return uint16(out)
	}
}

func SumOverLines(in string, f func(string) int) string {
	lines := strings.Split(in, "\n")
	sum := 0
	for _, line := range lines {
		sum += f(line)
	}
	return strconv.Itoa(sum)
}
