package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"valbaca.com/advent2015/day1"
)

const MAX_DAY = 1

func main() {
	day := getDay()
	input := readInputFile(day)
	fmt.Println(day1.Part2(input))
}

func getDay() int {
	if len(os.Args) < 2 {
		usageExit()
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > MAX_DAY {
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
