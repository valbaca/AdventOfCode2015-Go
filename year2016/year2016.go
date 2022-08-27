package year2016

import (
	"fmt"

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

// TODO: reduce the tedium of adding a new day
func ExecuteYear2016(day int, input string) {
	switch day {
	case 1:
		fmt.Println(day1.Part1(input), day1.Part2(input))
	case 2:
		fmt.Println(day2.Part1(input), day2.Part2(input))
	case 3:
		fmt.Println(day3.Part1(input), day3.Part2(input))
	case 4:
		fmt.Println(day4.Part1(input), day4.Part2(input))
	case 5:
		fmt.Println(day5.Part1(input), day5.Part2(input))
	case 6:
		fmt.Println(day6.Part1(input), day6.Part2(input))
	case 7:
		fmt.Println(day7.Part1(input), day7.Part2(input))
	case 8:
		fmt.Println(day8.Part1(input), day8.Part2(input))
	case 9:
		fmt.Println(day9.Part1(input), day9.Part2(input))
	case 10:
		fmt.Println(day10.Part1(input), day10.Part2(input))
	case 11:
		fmt.Println(day11.Part1(input), day11.Part2(input))
	case 12:
		fmt.Println(day12.Part1(input), day12.Part2(input))
	}
}
