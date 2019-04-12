package day19

import (
	"fmt"
	"strings"
)

func Part1(in, s string) string {
	reps := ParseInput(in)
	return fmt.Sprintf("%v", CountDistincts(reps, s))
}

func Part2(in, s string) string {
	reps := ParseInput(in)
	return fmt.Sprintf("%v", MinPath(reps, "e", s))
}

func ParseInput(in string) map[string][]string {
	reps := map[string][]string{}

	for _, line := range strings.Split(in, "\n") {
		fields := strings.Fields(line)
		old, new := fields[0], fields[2]
		if val, ok := reps[old]; ok {
			reps[old] = append(val, new)
		} else {
			reps[old] = []string{new}
		}
	}
	return reps
}

func CountDistincts(reps map[string][]string, s string) int {
	set := GetNewUniques(reps, s)
	return len(set)
}

func GetNewUniques(reps map[string][]string, s string) map[string]bool {
	set := map[string]bool{}
	for a, list := range reps {
		for _, b := range list {
			// a => b
			si := 0 // start index
			for si >= 0 && si < len(s) {
				idx := strings.Index(s[si:], a)
				if idx == -1 {
					break
				} else {
					idx += si
				}
				end := idx + len(a)
				new := s[0:idx] + b + s[end:]
				set[new] = true
				si = end
			}
		}
	}
	return set
}

func MinPath(reps map[string][]string, org, tgt string) int {
	return 0
}
