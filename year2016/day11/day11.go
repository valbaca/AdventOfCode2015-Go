package day11

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"gonum.org/v1/gonum/stat/combin"
	"valbaca.com/advent/elf"
)

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
*/

// A few globals purely for speed and simplicity
var NumFloors int
var NumItems int

func Part1(input string) string {
	state := parseInput(input)
	start := time.Now()
	minMoves := minMoves(state)
	fmt.Printf("\nPart 1 took %.6f\n", time.Since(start).Seconds())
	return strconv.Itoa(minMoves)
}

func Part2(input string) string {
	state := parseInput(input)
	state.floors[0].Add([]Item{
		{iso: "elerium", chip: false},
		{iso: "elerium", chip: true},
		{iso: "dilithium", chip: false},
		{iso: "dilithium", chip: true},
	})
	NumItems += 4
	start := time.Now()
	minMoves := minMoves(state)
	fmt.Printf("\nPart 2 took %.6f\n", time.Since(start).Seconds())
	return strconv.Itoa(minMoves)
}

func minMoves(init State) int {
	seen := map[string]bool{} // k: MinState.Hash, v: present?
	queue := []State{init}
	for len(queue) > 0 {
		state := &queue[0]
		queue = queue[1:]
		for _, optionState := range state.GenerateOptions() {
			minOptionState := optionState.MinState()
			minOptionStateHash := minOptionState.Hash()
			if _, hasSeen := seen[minOptionStateHash]; !hasSeen {
				seen[minOptionStateHash] = true
				// checking if solved sooner cuts runtime in half!
				if optionState.Solved() {
					return optionState.moves
				}
				queue = append(queue, optionState)
			}
		}
	}
	return -1
}

type MinState struct {
	// elev floor the elevator is on
	elev int
	// floorPairs: k: floor number, v: # of pairs on that floor
	floorPairs map[int]int
	// k1, floor of unpaired chip; k2, floor of unpaired generator, v: count of such unpaired (nearly always 1)
	unpaired map[int]map[int]int
}

func (ms *MinState) Hash() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%v#", ms.elev))
	for floorN := 0; floorN < NumFloors; floorN++ {
		sb.WriteString(fmt.Sprintf("f%vp%v", floorN, ms.floorPairs[floorN]))
		for genFloor, unpairCount := range ms.unpaired[floorN] {
			sb.WriteString(fmt.Sprintf("u%v->%v", unpairCount, genFloor))
		}
		sb.WriteRune(';')
	}
	return sb.String()
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
	iso  string
	chip bool // true if microchip; false if generator
}

func (s *State) Solved() bool {
	return s.elev == NumFloors-1 && len(s.floors[len(s.floors)-1].items) == NumItems
}

func (f Floor) PossibleMoves() [][]Item {
	// chain(combinations(f, 2), combinations(f, 1))
	numItems := len(f.items)
	if numItems == 0 {
		return nil
	}
	items := []Item{}
	for item := range f.items {
		items = append(items, item)
	}
	var takeTwo [][]int
	if numItems >= 2 {
		takeTwo = combin.Combinations(len(items), 2)
	}
	var takeOne [][]int
	if numItems >= 1 {
		takeOne = combin.Combinations(len(items), 1)
	}

	moves := make([][]Item, 0, len(takeTwo)+len(takeOne))
	for _, comb := range takeTwo {
		move := make([]Item, 2)
		move[0] = items[comb[0]] // normally would loop but there's only ever two
		move[1] = items[comb[1]]
		moves = append(moves, move)
	}
	for _, comb := range takeOne {
		move := make([]Item, 1)
		move[0] = items[comb[0]]
		moves = append(moves, move)
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

var dirs = [2]int{1, -1} // up and down

func (s *State) GenerateOptions() []State {
	options := []State{}
	possibleMoves := s.floors[s.elev].PossibleMoves()
	for _, move := range possibleMoves {
		for _, dir := range dirs {
			nextFloorN := s.elev + dir
			if nextFloorN < 0 || nextFloorN >= NumFloors {
				continue
			}
			// clone floors
			nextFloors := make([]Floor, len(s.floors))
			for i, floor := range s.floors {
				nextFloors[i] = floor.Clone()
			}
			// copy(nextFloors, s.floors)
			nextFloors[s.elev] = nextFloors[s.elev].Remove(move)
			if !nextFloors[s.elev].Safe() {
				continue
			}
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

func (s *State) MinState() *MinState {
	minState := MinState{s.elev, make(map[int]int), map[int]map[int]int{}}
	// start with chips, and find their generators, either on the same or diff floor
	for n, f := range s.floors {
		for item := range f.items {
			if item.chip {
			findGenerator:
				for gn, gf := range s.floors {
					for otherItem := range gf.items {
						if item.iso == otherItem.iso && !otherItem.chip {
							// found generator
							if n == gn {
								// on same floor, paired
								minState.floorPairs[n]++
								break findGenerator
							} else {
								// on different floors, unpaired
								if _, exists := minState.unpaired[n]; !exists {
									minState.unpaired[n] = make(map[int]int)
								}
								minState.unpaired[n][gn]++
							}
						}
					}
				}
			}
		}
	}
	return &minState
}

func parseInput(input string) State {
	lines := elf.Lines(input)
	NumFloors = len(lines)
	NumItems = 0
	floors := make([]Floor, 0, len(lines))
	for _, line := range lines {
		floor := Floor{make(map[Item]bool)}
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return !(unicode.IsLetter(r))
		})
		for i := len(fields) - 1; i > 0; i-- {
			field := fields[i]
			if field == "generator" {
				gen := Item{iso: fields[i-1], chip: false}
				floor.items[gen] = true // floor.items = append(floor.items, gen)
				NumItems++
			} else if field == "microchip" {
				iso := fields[i-2] // strings.TrimSuffix(, "-compatible")
				chip := Item{iso: iso, chip: true}
				floor.items[chip] = true //  floor.items = append(floor.items, chip)
				NumItems++
			}
		}
		floors = append(floors, floor)
	}
	return State{0, 0, floors}
}

// Nice to have, not necessary
func (s *State) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("m:%v\n", s.moves))
	for i := len(s.floors) - 1; i >= 0; i-- {
		floor := s.floors[i]
		currFloor := "  "
		if s.elev == i {
			currFloor = "->"
		}
		sb.WriteString(fmt.Sprintf("%v %v\n", currFloor, floor))
	}
	return sb.String()
}
