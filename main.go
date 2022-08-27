package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"valbaca.com/advent/elf"
	"valbaca.com/advent/year2015"
	"valbaca.com/advent/year2016"
)

var daysSolvedByYear = map[int]int{
	2015: 25,
	2016: 11, // UPDATE ME!
}

func main() {
	// Usage:
	// go run main.go               # no args => run all solutions
	// go run main.go [year] [day]  # run a specific year and day
	// go run main.go "latest"		# runs only the latest day's solution, tip: use watchexec

	// Uncomment to enable profiling
	// Then run:
	// $ go build -o pprofbin main.go && ./pprofbin latest && go tool pprof -http=":8000" pprofbin ./cpu.pprof
	/*
		f, perr := os.Create("cpu.pprof")
		if perr != nil {
			panic(perr)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	*/
	// TODO: use a flag instead of commenting/uncommenting

	if len(os.Args) >= 2 && os.Args[1] == "latest" {
		currYear := 0
		for y := range daysSolvedByYear {
			currYear = elf.Max(currYear, y)
		}
		execute(currYear, daysSolvedByYear[currYear])
		return
	}

	year, day := getYearAndDayFromArgs()
	if year > 0 && day > 0 {
		execute(year, day)
	} else {
		// execute all solutions from all years
		years := []int{2015, 2016}
		for _, y := range years {
			fmt.Printf("🎄 Year %v 🎄\n\n", y)
			for d := 1; d <= daysSolvedByYear[y]; d++ {
				execute(y, d)
			}
		}

	}
}

func execute(year, day int) {
	fmt.Printf("Year:%v Day:%v\n", year, day)
	input := readInputFile(year, day)
	input = strings.TrimSpace(input)
	start := time.Now()
	switch year {
	case 2015:
		year2015.ExecuteYear2015(day, input)
	case 2016:
		year2016.ExecuteYear2016(day, input)
	default:
		fmt.Printf("Invalid or incomplete year: %d\n", year)
	}
	elapsed := time.Since(start)
	fmt.Printf("took %.6fs\n\n", elapsed.Seconds())
}

func readInputFile(year, day int) string {
	name := fmt.Sprintf("./year%v/inputs/day%v.txt", year, day)
	out, err := os.ReadFile(name)
	if err != nil {
		// I'm tempted to remove this, but then immediately I forget to create an inputs file and I need it again!
		fmt.Fprintf(os.Stderr, "(No inputs file for year %v, day%v)\n", year, day)
	}
	return string(out)
}

func getYearAndDayFromArgs() (int, int) {
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
