package day12

/*
TIL: not a lot to this one. Just that `val, ok := map[key]` is pretty nice, made it easy to code getX and keep it fast
by avoiding an attempted atoi.
Using more type defs for more readable code.
*/
import (
	"strconv"
	"strings"

	"valbaca.com/advent/elf"
)

type Day12 struct{}

func (d Day12) Part1(input string) interface{} {
	cpu := NewCpu()
	cpu.exec(NewCmds(input))
	return strconv.Itoa(cpu["a"])
}

func (d Day12) Part2(input string) interface{} {
	cpu := NewCpu()
	cpu["c"] = 1
	cpu.exec(NewCmds(input))
	return strconv.Itoa(cpu["a"])
}

type Cpu map[string]int
type Cmd []string

func NewCmds(input string) []Cmd {
	lines := elf.Lines(input)
	cmds := make([]Cmd, len(lines))
	for i, line := range lines {
		cmds[i] = strings.Fields(line)
	}
	return cmds
}

func NewCpu() Cpu {
	return Cpu{"a": 0, "b": 0, "c": 0, "d": 0} // setting registers explicitly to make getX simple
}

func (cpu Cpu) exec(cmds []Cmd) {
	for i := 0; i >= 0 && i < len(cmds); i++ {
		cmd := cmds[i]
		switch cmd[0] {
		case "cpy":
			// cpy x y copies x (either an integer or the value of a register) into register y.
			x, y := cpu.getX(cmd[1]), cmd[2]
			cpu[y] = x
		case "inc":
			// inc x increases the value of register x by one.
			cpu[cmd[1]]++
		case "dec":
			// dec x decreases the value of register x by one.
			cpu[cmd[1]]--
		case "jnz":
			x, y := cpu.getX(cmd[1]), elf.UnsafeAtoi(cmd[2])
			if x != 0 {
				i += y
				i-- // account for end of loop i++
			}
			// jnz x y jumps to an instruction y away (positive means forward; negative means backward), but only if x is not zero.
		}
	}
}

func (cpu Cpu) getX(x string) int {
	if val, ok := cpu[x]; ok {
		return val
	}
	return elf.UnsafeAtoi(x)
}
