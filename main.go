package main

import (
	"bufio"
	"fmt"
	"os"

	"valbaca.com/advent2015/day1"
)

func main() {
	fmt.Println(day1.Part2(readFromStdin()))
}

func readFromStdin() string {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return in
}
