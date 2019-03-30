// utils provides short, unsafe versions of functions for use in hacking
package utils

import (
	"fmt"
	"strconv"
	"strings"
)

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
	lines := strings.Split(strings.TrimSpace(in), "\n")
	var sum int
	for _, line := range lines {
		sum += f(line)
	}
	return fmt.Sprintf("%v", sum)
}
