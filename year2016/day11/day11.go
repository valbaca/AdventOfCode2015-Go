package day11

/*
From when I did this one last, I remember it being one of THE hardest Advent problems.
I had to fall back to Java (my "native" programming language) and had to pull out all kinds of DS+A
like workstealing queues and backtracking, even then, part 1 took ~sec and part 2 took ~min
So going in...I already know I've got my work cut out for me...

Looking around online, a big optimization is storing a minimal "essence" of the problem in the `seen` set,
rather than storing the entire "true" state of the world (i.e. exclude the molecule names)

Took LOTS from this post:
https://eddmann.com/posts/advent-of-code-2016-day-11-radioisotope-thermoelectric-generators/

But still had plenty of work to do since Go doesn't make the following as trivial as Python:
- combinatorics
- tuples
- hashes
- copying: lists, maps, etc.

An optimization I also added was not going through all the floors searching for matching pairs. We have a map! Use it!

The final, awesome optimization was switching between using a string representation identifier of the minimum state
to using an integer hash code instead. The hash code logic isn't anything proven or verified, but it seems to work.

Finally, I got it down to around 0.6s! Compared to my Java solution (60s), that's nearly 100x faster.

Year:2016 Day:11

Part 1 took 0.115502s
Part 2 took 0.683379s
47 71
took 0.798930s
*/

import (
	"strings"
	"unicode"

	"gonum.org/v1/gonum/stat/combin"
	"valbaca.com/advent/elf"
)

type Day11 struct{}

func (d Day11) Part1(input string) interface{} {
	// start := time.Now()
	state, lastIso := parseInput(input)
	LastIso = lastIso
	minMoves := minMoves(state)
	// fmt.Printf("\nPart 1 took %.6fs", time.Since(start).Seconds())
	return minMoves
}

func (d Day11) Part2(input string) interface{} {
	// start := time.Now()
	state, lastIso := parseInput(input)
	// Add the new special elements
	state.floors[0].Add([]Item{
		{iso: lastIso, chip: false},
		{iso: lastIso, chip: true},
		{iso: lastIso + 1, chip: false},
		{iso: lastIso + 1, chip: true},
	})
	LastIso = lastIso + 2
	NumItems += 4
	minMoves := minMoves(state)
	// fmt.Printf("\nPart 2 took %.6fs\n", time.Since(start).Seconds())
	return minMoves
}

// A few globals purely for speed and simplicity
var NumFloors int
var NumItems int
var LastIso int

func minMoves(init State) int {
	seen := map[int]bool{} // k: MinState "hash", v: seen
	queue := []State{init}
	var state State
	for len(queue) > 0 {
		state, queue = queue[0], queue[1:]
		for _, option := range state.GenerateOptions() {
			optionHash := option.MinState().HashCode()
			if _, hasSeen := seen[optionHash]; !hasSeen {
				seen[optionHash] = true
				// checking if solved here (instead of top of queue-loop) cuts runtime in half!
				if option.Solved() {
					return option.moves
				}
				queue = append(queue, option)
			}
		}
	}
	return -1
}

type MinState struct {
	elev       int         // the floor the elevator is on
	floorPairs map[int]int // k: floor number, v: # of pairs on that floor
	// unpaired: k1, floor of unpaired chip; k2, floor of unpaired generator, v: count of such unpaired items (nearly always 1)
	unpaired map[int]map[int]int
}

func (ms MinState) HashCode() (hash int) {
	// random primes go!!!
	hash = 2
	hash += ms.elev * 5
	for n := 0; n < NumFloors; n++ {
		hash <<= 3
		hash += ms.floorPairs[n] * 7
		if len(ms.unpaired[n]) > 0 {
			for genFloor, unpairCount := range ms.unpaired[n] {
				hash += genFloor + 9*unpairCount
			}
		}
	}
	return
}

type State struct {
	moves  int
	elev   int
	floors []Floor
}

type Floor struct {
	items map[Item]bool
}

type Item struct {
	iso  int
	chip bool // true if microchip; false if generator
}

func (s State) Solved() bool {
	return s.elev == NumFloors-1 && len(s.floors[len(s.floors)-1].items) == NumItems
}

func (f Floor) PossibleMoves() [][]Item {
	numItems := len(f.items)
	if numItems == 0 {
		return nil
	}
	items := make([]Item, 0, numItems) // map keys into a list
	for item := range f.items {
		items = append(items, item)
	}
	var chooseTwo [][]int
	if numItems >= 2 {
		chooseTwo = combin.Combinations(numItems, 2)
	}
	var chooseOne [][]int
	if numItems >= 1 {
		chooseOne = combin.Combinations(numItems, 1)
	}

	moves := make([][]Item, 0, len(chooseTwo)+len(chooseOne))
	for _, comb := range chooseTwo {
		moves = append(moves, []Item{items[comb[0]], items[comb[1]]})
	}
	for _, comb := range chooseOne {
		moves = append(moves, []Item{items[comb[0]]})
	}

	return moves
}

func (f Floor) Clone() Floor {
	orig := f.items
	clone := make(map[Item]bool, len(orig))
	for k, v := range orig {
		clone[k] = v
	}
	return Floor{clone}
}

func (f Floor) Add(items []Item) Floor {
	for _, item := range items {
		f.items[item] = true
	}
	return f
}

func (f Floor) Remove(items []Item) Floor {
	for _, item := range items {
		delete(f.items, item)
	}
	return f
}

func (f Floor) Safe() bool {
	hasUnpairedChip, hasUnpairedGen := false, false
	for item, _ := range f.items {
		if !f.items[Item{iso: item.iso, chip: !item.chip}] {
			if item.chip {
				hasUnpairedChip = true
			} else {
				hasUnpairedGen = true
			}
		}
	}
	return !(hasUnpairedGen && hasUnpairedChip)
}

func (s State) GenerateOptions() []State {
	options := []State{}
	possibleMoves := s.floors[s.elev].PossibleMoves()
	for _, move := range possibleMoves {
		for _, dir := range [2]int{1, -1} { // go up then down
			nextFloorN := s.elev + dir
			if nextFloorN < 0 || nextFloorN >= NumFloors {
				continue
			}
			nextFloors := make([]Floor, len(s.floors))
			// can shallow copy most floors, but we need to clone the floors modified
			copy(nextFloors, s.floors)
			nextFloors[s.elev] = s.floors[s.elev].Clone()
			nextFloors[s.elev] = nextFloors[s.elev].Remove(move)
			if !nextFloors[s.elev].Safe() {
				continue
			}

			nextFloors[nextFloorN] = s.floors[nextFloorN].Clone()
			nextFloors[nextFloorN] = nextFloors[nextFloorN].Add(move)
			if !nextFloors[nextFloorN].Safe() {
				continue
			}
			option := State{moves: s.moves + 1, elev: nextFloorN, floors: nextFloors}
			options = append(options, option)
		}
	}
	return options
}

func (s State) MinState() MinState {
	minState := MinState{s.elev, make(map[int]int), map[int]map[int]int{}}
	// start with chips, and find their generators, either on the same or diff floor
	for n, f := range s.floors {
		for item := range f.items {
			if item.chip {
				gen := Item{iso: item.iso, chip: false}
				for gn, gf := range s.floors {
					if _, ok := gf.items[gen]; ok {
						// found generator
						if n == gn {
							// found generator on same floor => paired
							minState.floorPairs[n]++
							break
						} else {
							// found generator on different floors => unpaired
							if _, exists := minState.unpaired[n]; !exists {
								minState.unpaired[n] = make(map[int]int)
							}
							minState.unpaired[n][gn]++
							break
						}
					}
				}
			}
		}
	}
	return minState
}

func parseInput(input string) (state State, newIso int) {
	lines := elf.Lines(input)
	NumFloors = len(lines)
	NumItems = 0
	nameToIso := map[string]int{}
	floors := make([]Floor, 0, len(lines))
	for _, line := range lines {
		floor := Floor{make(map[Item]bool)}
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return !(unicode.IsLetter(r))
		})
		for i := len(fields) - 1; i > 0; i-- {
			field := fields[i]
			var name string
			var chip bool
			if field == "generator" {
				name = fields[i-1]
			} else if field == "microchip" {
				name = fields[i-2]
				chip = true
			}
			if name != "" {
				iso, seen := nameToIso[name]
				if !seen {
					iso, nameToIso[name] = newIso, newIso
					newIso++
				}
				item := Item{iso, chip}
				floor.items[item] = true
				NumItems++
			}
		}
		floors = append(floors, floor)
	}
	state.floors = floors
	return
}
