// Package day14
// TIL: not a lot in this one. I guess that Go is really, really fast?
package day14

import (
	"strconv"
	"strings"

	"valbaca.com/advent2015/utils"
)

type raindeer struct {
	speed   int // in km/sec
	stamina int // # of secs raindeer can fly at speed
	sleep   int // secs needed to rest
	period  int // time, in sec, for a full fly-rest cycle (stamina + sleep)
	dist    int // dist traveled in one period (speed * stamina)
	points  int
}

func Part1(in string, stop int) string {
	deers := ParseInput(in)
	dists := CalcDists(deers, stop)
	return strconv.Itoa(utils.Max(dists...))
}

func Part2(in string, stop int) string {
	deers := ParseInput(in)
	for i := 1; i <= stop; i++ {
		dists := CalcDists(deers, i)
		max, maxIndex := dists[0], 0
		for j := 1; j < len(dists); j++ {
			if dists[j] > max {
				max = dists[j]
				maxIndex = j
			}
		}
		(&deers[maxIndex]).incPoints()
	}
	maxPoints := utils.MinInt
	for _, deer := range deers {
		if deer.points > maxPoints {
			maxPoints = deer.points
		}
	}
	return strconv.Itoa(maxPoints)
}

func ParseInput(in string) []raindeer {
	out := []raindeer{}
	for _, line := range strings.Split(in, "\n") {
		out = append(out, ParseLine(line))
	}
	return out
}

func ParseLine(line string) raindeer {
	// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	// 0     1   2   3  4    5   6  7        8   9    10   11   12  13  14
	sp := strings.Split(line, " ")
	speed := utils.Atoi(sp[3])
	stamina := utils.Atoi(sp[6])
	sleep := utils.Atoi(sp[13])
	return raindeer{speed, stamina, sleep, stamina + sleep, speed * stamina, 0}
}

func CalcDists(rdeers []raindeer, stop int) []int {
	out := []int{}
	for _, deer := range rdeers {
		out = append(out, deer.CalcDist(stop))
	}
	return out
}

func (rd raindeer) CalcDist(time int) int {
	// |---|----|
	dist := 0
	if time >= rd.period {
		fullPeriods := time / rd.period
		dist = fullPeriods * (rd.dist)
		time = time % rd.period
	}
	if time > rd.stamina {
		time = rd.stamina
	}
	return dist + time*rd.speed
}

func (rd *raindeer) incPoints() {
	rd.points++
}
