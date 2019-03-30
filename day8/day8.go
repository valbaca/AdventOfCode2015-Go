package day8

import (
	"strings"

	"valbaca.com/advent2015/utils"
)

// between 1222, 1344
// guess 1: 1507
func Part1(in string) string {
	return utils.SumOverLines(in, GetOverhead)
}

func Part2(in string) string {
	return utils.SumOverLines(in, GetAddedOverhead)
}

func GetOverhead(s string) int {
	oh := 2
	sp := strings.Split(strings.TrimSpace(s), "")
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
func GetAddedOverhead(s string) int {
	oh := 2
	sp := strings.Split(strings.TrimSpace(s), "")
	for i := 0; i < len(sp); i++ {
		switch sp[i] {
		case `"`, `\`:
			oh++
		}
	}
	return oh
}
