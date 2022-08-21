// Package day12
// TIL: hand parsing JSON sucks
// TIL: regexs in go are very fast but limited b/c no back references
package day12

import (
	"fmt"
	"regexp"
	"valbaca.com/advent/elf"
)

func Part1(in string) string {
	return sumString(in)
}

// 68214 too high
func Part2(in string) string {
	//return sumString(filterRed(strings.TrimSpace(in)))
	return sumNoRed(in)
}

func sumString(s string) string {
	buf := ""
	sum := 0
	for i := 0; i < len(s); i++ {
		si := s[i]
		if (si >= '0' && si <= '9') || si == '-' {
			buf = fmt.Sprintf("%s%c", buf, si)
		} else if len(buf) != 0 {
			sum += elf.UnsafeAtoi(buf)
			buf = ""
		}
	}
	if len(buf) != 0 {
		sum += elf.UnsafeAtoi(buf)
	}
	return fmt.Sprintf("%d", sum)
}

// This is really, really gross looking but it gets it done!
func sumNoRed(s string) string {
	buf := ""
	sum := 0
	for i := 0; i < len(s); i++ {
		si := s[i]
		if (si >= '0' && si <= '9') || si == '-' {
			buf = fmt.Sprintf("%s%c", buf, si)
		} else if len(buf) != 0 {
			sum += elf.UnsafeAtoi(buf)
			buf = ""
		}
		if si == '{' {
			openBrackets := 1
			var closeIndex int
			for j := i + 1; j < len(s); j++ {
				if s[j] == '{' {
					openBrackets++
				}
				if s[j] == '}' {
					openBrackets--
					if openBrackets == 0 {
						closeIndex = j
						break
					}
				}
			}
			if closeIndex == 0 {
				panic("uh oh")
			}
			if closeIndex-i > 1 {
				sum += elf.UnsafeAtoi(sumNoRed(s[i+1 : closeIndex]))
			}
			i = closeIndex
		}
		if si == ':' && len(s)-i >= 6 && s[i:i+6] == `:"red"` {
			//fmt.Printf("Sum of %s is 0 (red found!)\n", s) // Debug
			return "0"
		}
	}
	if len(buf) != 0 {
		sum += elf.UnsafeAtoi(buf)
	}
	out := fmt.Sprintf("%d", sum)
	//fmt.Printf("Sum of %s is %s\n", s, out) // Debug
	return out
}

func filterRed(s string) string {
	re := regexp.MustCompile(`(\{.+?:"red".+?\})`)
	out := re.ReplaceAllString(s, "")
	fmt.Println(out)
	return out
}
