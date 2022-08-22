package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"valbaca.com/advent/year2015"
	"valbaca.com/advent/year2016"
)

func main() {
	// Usage:
	// go run main.go               # no args => run all solutions
	// go run main.go [year] [day]  # run a specific year and day
	// go run main.go "latest"		# runs only the latest day's solution, tip: use watchexec

	if len(os.Args) >= 2 && os.Args[1] == "latest" {
		currYear := 2015
		execute(currYear, daysSolvedByYear[currYear])
		return
	}

	year, day := getYearAndDay()
	if year > 0 && day > 0 {
		execute(year, day)
	} else {
		// execute all solutions from all years
		years := []int{2015, 2016}
		for _, y := range years {
			fmt.Printf("🎄 Year %v 🎄\n\n", y)
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
		year2015.ExecuteYear2015(day, input)
	case 2016:
		year2016.ExecuteYear2016(day, input)
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

var daysSolvedByYear = map[int]int{
	2015: 25,
	2016: 1,
	// UPDATE ME!
}
