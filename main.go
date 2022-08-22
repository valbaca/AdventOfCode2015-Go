package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"valbaca.com/advent/year2015/day1"
	"valbaca.com/advent/year2015/day10"
	"valbaca.com/advent/year2015/day11"
	"valbaca.com/advent/year2015/day12"
	"valbaca.com/advent/year2015/day13"
	"valbaca.com/advent/year2015/day14"
	"valbaca.com/advent/year2015/day15"
	"valbaca.com/advent/year2015/day16"
	"valbaca.com/advent/year2015/day17"
	"valbaca.com/advent/year2015/day18"
	"valbaca.com/advent/year2015/day19"
	"valbaca.com/advent/year2015/day2"
	"valbaca.com/advent/year2015/day20"
	"valbaca.com/advent/year2015/day21"
	"valbaca.com/advent/year2015/day3"
	"valbaca.com/advent/year2015/day4"
	"valbaca.com/advent/year2015/day5"
	"valbaca.com/advent/year2015/day6"
	"valbaca.com/advent/year2015/day7"
	"valbaca.com/advent/year2015/day8"
	"valbaca.com/advent/year2015/day9"
	y16d1 "valbaca.com/advent/year2016/day1"
)

var daysSolvedByYear = map[int]int{
	2015: 21,
	2016: 1,
}

func main() {
	// Usage:
	// go run main.go               # no args => run all solutions
	// go run main.go [year] [day]  # run a specific year and day
	// go run main.go "latest"		# runs only the latest day's solution, tip: use watchexec

	if len(os.Args) >= 2 && os.Args[1] == "latest" {
		execute(2015, 21)
		return
	}

	year, day := getYearAndDay()
	if year > 0 && day > 0 {
		execute(year, day)
	} else {
		// execute all solutions from all years
		years := []int{2015, 2016}
		for _, y := range years {
			for d := 1; d <= daysSolvedByYear[y]; d++ {
				start := time.Now()

				fmt.Printf("Day %d: ", d)
				execute(y, d)

				elapsed := time.Since(start)
				fmt.Printf("took %.6fs\n\n", elapsed.Seconds())
			}
		}

	}
}

func execute(year, day int) {
	input := readInputFile(year, day)
	input = strings.TrimSpace(input)
	switch year {
	case 2015:
		executeYear2015(day, input)
	case 2016:
		executeYear2016(day, input)
	default:
		fmt.Printf("Invalid or incomplete year: %d\n", year)
	}
}

func getYearAndDay() (int, int) {
	if len(os.Args) < 2 {
		return 0, 0
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, 0
	}
	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return 0, 0
	}
	return year, day
}

func readInputFile(year, day int) string {
	name := fmt.Sprintf("./year%v/inputs/day%v.txt", year, day)
	out, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "(No input file for year %v, day%v)\n", year, day)
	}
	return string(out)
}

func executeYear2015(day int, input string) {
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
		fmt.Println(day9.BothParts(input))
	case 10:
		fmt.Println(day10.Part1(input), day10.Part2(input))
	case 11:
		fmt.Println(day11.Part1(input), day11.Part2(input))
	case 12:
		fmt.Println(day12.Part1(input), day12.Part2(input))
	case 13:
		fmt.Println(day13.Part1(input), day13.Part2(input))
	case 14:
		fmt.Println(day14.Part1(input, 2503), day14.Part2(input, 2503))
	case 15:
		fmt.Println(day15.Part1(input), day15.Part2(input))
	case 16:
		fmt.Println(day16.Part1(input), day16.Part2(input))
	case 17:
		fmt.Println(day17.Part1(input, 150), day17.Part2(input, 150))
	case 18:
		fmt.Println(day18.Part1(input, 100, 100), day18.Part2(input, 100, 100))
	case 19:
		molecule := "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"
		fmt.Println(day19.Part1(input, molecule), day19.Part2(input, molecule))
	case 20:
		fmt.Println(day20.Part1(33100000), day20.Part2(33100000))
	case 21:
		fmt.Println(day21.Part1(input, 100, 103, 9, 2), day21.Part2(input, 100, 103, 9, 2))
	}
}

func executeYear2016(day int, input string) {
	switch day {
	case 1:
		fmt.Println(y16d1.Howdy())
	}
}
