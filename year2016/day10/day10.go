package day10

/*
TIL: Trying out VS Code to wean off of GoLand (just in case I want to stick to purely free tools in the future)
*/
import (
	"sort"
	"strconv"
	"strings"

	"valbaca.com/advent/elf"
)

var atoi = elf.UnsafeAtoi
var output string

func Part1(input string) string {
	lines := elf.Lines(input)
	bots = make(map[int]*Bot)
	outputs = make(map[int]int)
	for _, line := range lines {
		readLine(line)
	}
	return output
}

type Bot struct {
	id   int
	vals []int
	inst string
}

var bots map[int]*Bot
var outputs map[int]int

func readLine(line string) {
	fields := strings.Fields(line)
	if fields[0] == "value" {
		// value 5 goes to bot 2
		val, target, id := atoi(fields[1]), fields[4], atoi(fields[5])
		if target != "bot" {
			panic("Only expected bot")
		}
		getBot(id).sendValue(val)

	} else if fields[0] == "bot" {
		// bot 2 gives low to bot 1 and high to bot 0
		id := atoi(fields[1])
		getBot(id).sendInst(line)
	}
}

func getBot(id int) *Bot {
	if bot, ok := bots[id]; ok {
		return bot
	}
	bot := new(Bot)
	bot.id = id
	bots[id] = bot
	return bot
}

func (b *Bot) sendValue(value int) {
	b.vals = append(b.vals, value)
	b.tryExec()
}

func (b *Bot) sendInst(s string) {
	b.inst = s
	b.tryExec()
}

func (b *Bot) tryExec() {
	if b.inst == "" || len(b.vals) != 2 {
		return
	}
	sort.Slice(b.vals, func(i, j int) bool { return b.vals[i] < b.vals[j] })
	lowValue, highValue := b.vals[0], b.vals[1]
	if lowValue == 17 && highValue == 61 {
		output = strconv.Itoa(b.id)
	}
	// bot 1 gives low to output 1 and high to bot 0
	// 0   1 2     3   4  5      6 7   8    9  10  11
	fields := strings.Fields(b.inst)
	lowType, lowId, highType, highId := fields[5], atoi(fields[6]), fields[10], atoi(fields[11])
	sendValue(lowType, lowId, lowValue)
	sendValue(highType, highId, highValue)
}

func sendValue(targetType string, id int, value int) {
	if targetType == "output" {
		outputs[id] = value
	} else {
		getBot(id).sendValue(value)
	}
}

func Part2(input string) string { // assumes Part1 has run
	return strconv.Itoa(outputs[0] * outputs[1] * outputs[2])
}
