// Package day8
// TIL: for these problems, trimming the input is really key
package day8

import (
	"strconv"
	"strings"
)

// between 1222, 1344
// guess 1: 1507
func Part1(in string) string {
	return sumOverLines(in, getOverhead)
}

func Part2(in string) string {
	return sumOverLines(in, getAddedOverhead)
}

func getOverhead(s string) int {
	oh := 2
	sp := strings.Split(s, "")
	for i := 1; i < len(sp)-1; i++ {
		if sp[i] == `\` {
			if sp[i+1] == `\` {
				oh++
			} else if sp[i+1] == `"` {
				oh++
			} else if sp[i+1] == `x` {
				oh += 3
			}
			i++ // skip ahead to prevent double counts
		}
	}
	return oh
}

// [0,2719)
func getAddedOverhead(s string) int {
	oh := 2
	sp := strings.Split(s, "")
	for i := 0; i < len(sp); i++ {
		switch sp[i] {
		case `"`, `\`:
			oh++
		}
	}
	return oh
}

func sumOverLines(in string, f func(string) int) string {
	lines := strings.Split(in, "\n")
	sum := 0
	for _, line := range lines {
		sum += f(line)
	}
	return strconv.Itoa(sum)
}
