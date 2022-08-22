package day23

/*
TIL: I do miss my `sep` function from Clojure, got stuck on a single comma for too long.
`sep` was basically the ideal scanner function, split on whitespace, colons and commas and safely attempted to convert
strings to ints. That's the kind of thing dynamic languages can do that static languages just really cannot (at least
not easily).
*/
import (
	"strconv"
	"strings"
	"valbaca.com/advent/elf"
)

type Command struct {
	op   string
	args []string
}

type Comp struct {
	commands  []Command
	currLine  int
	registers map[string]int
}

func Part1(input string) string {
	commands := parseInput(input)
	comp := Comp{commands, 0, map[string]int{"a": 0, "b": 0}}
	comp.exec()
	return strconv.Itoa(comp.registers["b"])
}

func Part2(input string) string {
	commands := parseInput(input)
	comp := Comp{commands, 0, map[string]int{"a": 1, "b": 0}}
	comp.exec()
	return strconv.Itoa(comp.registers["b"])
}
func parseInput(input string) []Command {
	lines := strings.Split(input, "\n")
	commands := make([]Command, len(lines))
	for i := 0; i < len(lines); i++ {
		splits := elf.SplitWords(lines[i])
		for j := 0; j < len(splits); j++ {
			splits[j] = strings.TrimRight(splits[j], ",") // stupid comma
		}
		commands[i] = Command{splits[0], splits[1:]}
	}
	return commands
}

func (c *Comp) exec() {
	for c.currLine >= 0 && c.currLine < len(c.commands) {
		jumped := false
		cmd := c.commands[c.currLine]
		arg := cmd.args[0]
		switch cmd.op {
		case "hlf":
			c.registers[arg] /= 2
		case "tpl":
			c.registers[arg] *= 3
		case "inc":
			c.registers[arg]++
		case "jmp":
			jumped = c.jump(elf.UnsafeAtoi(arg))
		case "jie":
			if c.registers[arg]%2 == 0 {
				jumped = c.jump(elf.UnsafeAtoi(cmd.args[1]))
			}
		case "jio":
			if c.registers[arg] == 1 {
				jumped = c.jump(elf.UnsafeAtoi(cmd.args[1]))
			}
		}
		if !jumped {
			c.next()
		}
	}
}

func (c *Comp) jump(n int) bool {
	c.currLine += n
	return true
}

func (c *Comp) next() {
	c.jump(1)
}
