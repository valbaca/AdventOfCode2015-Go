package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"valbaca.com/advent2015/day1"
	"valbaca.com/advent2015/day2"
	"valbaca.com/advent2015/day3"
)

func main() {
	day := getDay()
	// TODO don't read whole file in memory
	input := readInputFile(day)
	switch day {
	case 1:
		fmt.Println(day1.Part1(input), day1.Part2(input))
	case 2:
		fmt.Println(day2.Part1(input), day2.Part2(input))
	case 3:
		fmt.Println(day3.Part1(input), day3.Part2(input))
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
