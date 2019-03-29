package day7

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"valbaca.com/advent2015/utils"
)

func Part1(in, key string) string {
	wires := ReadInWires(in)
	val, err := GetVal(key, wires)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(val)
}

func Part2(in, key string) string {
	wires := ReadInWires(in)
	origA, err := GetVal(key, wires)
	if err != nil {
		panic(err)
	}
	wires = ReadInWires(in)
	wires["b"] = Wire{"SET", nil, "b", origA} // Override before re-running GetVal
	val, err := GetVal(key, wires)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(val)
}

func ReadInWires(in string) map[string]Wire {
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	wires := make(map[string]Wire) //, 26*26)
	for _, line := range lines {
		wire, err := ReadLine(line)
		if err != nil {
			fmt.Println(err)
		}
		if wire != nil {
			wires[wire.out] = *wire
		}
	}
	return wires
}

type Wire struct {
	op  string   // "AND" etc.
	ins []string // labels of the inputs
	out string   // label of the output
	arg uint16   // integer argument, used for SET, LSHIFT, and RSHIFT
}

func ReadLine(s string) (*Wire, error) {
	sp := strings.Split(s, " ")
	if strings.Contains(s, "AND") {
		return ReadAndOr(sp)
	} else if strings.Contains(s, "OR") {
		return ReadAndOr(sp)
	} else if strings.Contains(s, "SHIFT") {
		return ReadShift(sp)
	} else if strings.Contains(s, "NOT") {
		return ReadNot(sp)
	}
	// value assignment, easy!
	return ReadSet(sp)
}

func ReadNot(s []string) (*Wire, error) {
	// NOT x -> h
	if len(s) < 4 {
		return nil, errors.New("NOT couldn't parse")
	}
	op, x, out := s[0], s[1], s[3]
	return &Wire{op, []string{x}, out, 0}, nil
}
func ReadShift(s []string) (*Wire, error) {
	// x LSHIFT 2 -> f
	if len(s) < 5 {
		return nil, errors.New("SHIFT couldn't parse")
	}
	x, op, arg, out := s[0], s[1], utils.AtoUint16(s[2]), s[4]
	return &Wire{op, []string{x}, out, arg}, nil
}

func ReadAndOr(s []string) (*Wire, error) {
	// x AND y -> out
	if len(s) != 5 {
		return nil, errors.New("AND/OR couldn't parse")
	}
	x, op, y, out := s[0], s[1], s[2], s[4]
	return &Wire{op, []string{x, y}, out, 0}, nil
}

func ReadSet(s []string) (*Wire, error) {
	op := "SET"
	if len(s) != 3 {
		return nil, errors.New("SET couldn't parse")
	}
	lh, rh := s[0], s[2]
	if lv, err := strconv.ParseUint(lh, 10, 16); err == nil {
		return &Wire{op, nil, rh, uint16(lv)}, nil
	}
	return &Wire{op, []string{lh}, rh, 0}, nil
}

func GetVal(key string, wires map[string]Wire) (uint16, error) {
	wire, ok := wires[key]
	if !ok {
		// one last try...
		if v, err := strconv.ParseUint(key, 10, 16); err == nil {
			return uint16(v), nil
		}
		return 0, errors.New("missing wire:" + key)
	}
	switch wire.op {
	case "SET":
		if wire.ins == nil {
			return wire.arg, nil
		}
		out, err := GetVal(wire.ins[0], wires)
		if err != nil {
			return 0, err
		}
		//wires[wire.out] = Wire{"SET", nil, wire.out, out} // optimization?
		return out, nil
	case "AND":
		x, err := GetVal(wire.ins[0], wires)
		if err != nil {
			return 0, err
		}
		y, err := GetVal(wire.ins[1], wires)
		if err != nil {
			return 0, err
		}
		out := x & y
		//wires[wire.out] = Wire{"SET", nil, wire.out, out} // optimization?
		return out, nil
	case "OR":
		x, err := GetVal(wire.ins[0], wires)
		if err != nil {
			return 0, err
		}
		y, err := GetVal(wire.ins[1], wires)
		if err != nil {
			return 0, err
		}
		out := x | y
		wires[wire.out] = Wire{"SET", nil, wire.out, out} // optimization?
		return out, nil
	case "LSHIFT":
		x, err := GetVal(wire.ins[0], wires)
		if err != nil {
			return 0, err
		}
		out := x << wire.arg
		//wires[wire.out] = Wire{"SET", nil, wire.out, out} // optimization?
		return out, nil
	case "RSHIFT":
		x, err := GetVal(wire.ins[0], wires)
		if err != nil {
			return 0, err
		}
		out := x >> wire.arg
		//wires[wire.out] = Wire{"SET", nil, wire.out, out} // optimization?
		return out, nil
	case "NOT":
		x, err := GetVal(wire.ins[0], wires)
		if err != nil {
			return 0, err
		}
		out := ^x
		//wires[wire.out] = Wire{"SET", nil, wire.out, out} // optimization?
		return out, nil
	}
	return 0, errors.New("not done")
}
