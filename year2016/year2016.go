package year2016

import (
	"valbaca.com/advent/elf"
	"valbaca.com/advent/year2016/day1"
	"valbaca.com/advent/year2016/day10"
	"valbaca.com/advent/year2016/day11"
	"valbaca.com/advent/year2016/day12"
	"valbaca.com/advent/year2016/day2"
	"valbaca.com/advent/year2016/day3"
	"valbaca.com/advent/year2016/day4"
	"valbaca.com/advent/year2016/day5"
	"valbaca.com/advent/year2016/day6"
	"valbaca.com/advent/year2016/day7"
	"valbaca.com/advent/year2016/day8"
	"valbaca.com/advent/year2016/day9"
)

var Year2016 []elf.Day = []elf.Day{
	nil, // 0
	day1.Day1{},
	day2.Day2{},
	day3.Day3{},
	day4.Day4{},
	day5.Day5{},
	day6.Day6{},
	day7.Day7{},
	day8.Day8{},
	day9.Day9{},
	day10.Day10{},
	day11.Day11{},
	day12.Day12{},
}

// TODO: reduce the tedium of adding a new day
func ExecuteYear2016(d int, input string) {
	// year2016 := []elf.Day{
	// 	nil, // 0
	// 	day1.Day1{},
	// 	day2.Day2{},
	// 	day3.Day3{},
	// 	day4.Day4{},
	// 	day5.Day5{},
	// 	day6.Day6{},
	// 	day7.Day7{},
	// 	day8.Day8{},
	// 	day9.Day9{},
	// 	day10.Day10{},
	// 	day11.Day11{},
	// 	day12.Day12{},
	// }
	if d <= 0 || d >= len(Year2016) {
		panic("Invalid day")
	}
	elf.ExecDay(Year2016[d], input)
}
