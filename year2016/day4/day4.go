package day4

/*
TIL: needing to copy in a priority queue implementation and needing to have it be hard-coded to one specific use case is
...not ideal? Tons easier in Python and even Java
*/
import (
	"container/heap"
	"strconv"
	"strings"
	"unicode"

	"valbaca.com/advent/elf"
)

type Day4 struct{}

func (d Day4) Part1(input string) interface{} {
	rooms := parse(input)
	ans := 0
	for _, room := range rooms {
		if room.isValid() {
			ans += room.sectorId
		}
	}
	return ans
}

type room struct {
	line     string
	name     string
	sectorId int
	checksum string
}

func parse(input string) []room {
	lines := strings.Split(input, "\n")
	var rooms []room
	for _, line := range lines {
		rooms = append(rooms, newRoom(line))
	}
	return rooms
}

func newRoom(line string) room {
	nameBuilder := strings.Builder{}
	endNameIdx := -1
	for i, r := range line {
		if unicode.IsDigit(r) {
			endNameIdx = i
			break
		} else if r == '-' {
			continue
		}
		nameBuilder.WriteRune(r)
	}
	name := nameBuilder.String()
	sectorIdStringBuilder := strings.Builder{}
	for _, r := range line[endNameIdx:] {
		if !unicode.IsDigit(r) {
			break
		}
		sectorIdStringBuilder.WriteRune(r)
	}
	sectorId := elf.UnsafeAtoi(sectorIdStringBuilder.String())
	checksum := line[len(line)-6 : len(line)-1]
	return room{line, name, sectorId, checksum}
}

func (r room) isValid() bool {
	calculatedChecksum := r.calcChecksum()
	return r.checksum == calculatedChecksum
}

func (r room) calcChecksum() string {
	counts := map[rune]int{}
	for _, char := range r.name {
		counts[char]++
	}
	pq := make(elf.PriorityQueue, len(counts))
	i := 0
	for char, count := range counts {
		pq[i] = &elf.Item{Value: string(char), Priority: count, Index: i}
		i++
	}
	heap.Init(&pq)
	sb := strings.Builder{}
	for j := 0; j < 5; j++ {
		item := heap.Pop(&pq).(*elf.Item)
		sb.WriteString(item.Value)
	}
	return sb.String()
}

func (d Day4) Part2(input string) interface{} {
	rooms := parse(input)
	for _, room := range rooms {
		if room.isValid() {
			cipher := Rotate(room.name, room.sectorId)
			//println(cipher) // only by printing them out and manually scanning do we know to look for "pole"
			if strings.Contains(cipher, "pole") {
				return strconv.Itoa(room.sectorId)
			}
		}
	}
	return "FAIL"
}

func Rotate(input string, n int) string {
	sb := strings.Builder{}
	for _, ch := range input {
		alphIdx := ((int(ch) - int('a')) + n) % 26
		sb.WriteRune(rune(alphIdx + 'a'))
	}
	return sb.String()
}
