# Advent of Code in Go

Solutions to Advent of Code problems in Go.

## Running

```shell
go run main.go               # runs all solutions
go run main.go [year] [day]  # runs a particular day
go run main.go latest        # runs the last day solved
```

(Nearly) Every day contains a "TIL" (Today I Learned) message that explains something about Go that I learned.

Looking back, the most common "TIL" themes are:

1. Wow. Go is *FAST*. So fast it can make brute-force solutions viable. Similarly, if any code is ever taking longer than a few seconds, you know a better solution exists (as opposed to other languages where `O(n)` solutions can still execute slowly).
   1. Compared to Python, between 10x and 100x as fast, usually 25-50x as fast. Solutions typically take ~33-50% more lines than their Python equivalent.
2. Conversion between types is cumbersome. A lot of Advent of Code is parsing various ad-hoc formats for numbers and text commands.
   1. I miss simple splitters and scanners and converters from Python and Clojure
3. bytes.Buffer appending is much faster than strings (as it is in most languages).
4. Go's built-in Regex is fast but limited; no back-references.
5. Finding the unit testing boundary is tricky. Nearly every question has some sort of "parse input into structs and execute logic." Sometimes the parse logic is a majority of the code, other times it's trivial, so unit tests may be the actual text input, or it could be using the intermediate structs.
6. A few of the core functionalities have foot-guns: slices and `:=` vs `=`
7. I really miss niceties like tuples, `filter/map/flatMap/reduce`, and function overloading and default args.


## Tips

Use [watchexec](https://watchexec.github.io/) to automatically re-run the code on file save:
```shell
$ watchexec go run main.go [year] [day]
```

## Structure

In an attempt to keep the structure sane, going with the following:

```
go.mod         # declares everything under 'valbaca.com/advent' package (omitting this package prefix below for brevity)
main.go        # package main; the main runner for all solutions
               # main runner allows running either: latest, all, or one particular day
               # main runner handles reading input from file, given its put in the right place
               # of `year{y}/input/day{d}.txt`, e.g. `year2015/input/day1.txt`
elf/elf.go     # package elf, little helper functions
year2015/
  year2015.go  # year2015/year2015.go has package year2015 and has ExecuteYear2015 which executes the individual days
  day1/
    day1.go    # year2015/day1/day1.go has package year2015/day1
    day1_test.go  # testing, as desired
  day2/        # each day gets a separate folder and file...TODO, simplify this to just a file?
    day2.go
  ...
  day25/
    day25.go
  inputs/      # all year2015 input files go here, easy to add. The 
    day1.txt
    day2.txt
year2016/
  year2016.go  # and so on...
```