package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"valbaca.com/advent2015/day1"
	"valbaca.com/advent2015/day10"
	"valbaca.com/advent2015/day11"
	"valbaca.com/advent2015/day12"
	"valbaca.com/advent2015/day2"
	"valbaca.com/advent2015/day3"
	"valbaca.com/advent2015/day4"
	"valbaca.com/advent2015/day5"
	"valbaca.com/advent2015/day6"
	"valbaca.com/advent2015/day7"
	"valbaca.com/advent2015/day8"
	"valbaca.com/advent2015/day9"
)

func main() {
	day := getDay()
	// TODO don't read whole file in memory
	// TODO trim the input so each program doesn't have to
	input := readInputFile(day)
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
		fmt.Println(day7.Part1(input, "a"), day7.Part2(input, "a"))
	case 8:
		fmt.Println(day8.Part1(input), day8.Part2(input))
	case 9:
		fmt.Println(day9.Part1(input))
	case 10:
		fmt.Println(day10.Part1(input), day10.Part2(input))
	case 11:
		fmt.Println(day11.Part1(input), day11.Part2(input))
	case 12:
		fmt.Println(day12.Part1(input), day12.Part2(input))
	}
}

func getDay() int {
	if len(os.Args) < 2 {
		usageExit()
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usageExit()
	}
	return day
}

func usageExit() {
	fmt.Fprintln(os.Stderr, "Usage: [day]\nday must be int")
	os.Exit(1)
}

func readInputFile(day int) string {
	name := fmt.Sprintf("./inputs/day%v.txt", day)
	out, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(out)
}
