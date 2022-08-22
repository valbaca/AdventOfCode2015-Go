package day25

import "fmt"

const (
	targetRow = 2978
	targetCol = 3083
)

type Pos struct {
	r, c, val int64
}

func Part1() string {
	p := Pos{1, 1, 20151125}
	for {
		if p.r == targetRow && p.c == targetCol {
			return fmt.Sprintf("%v", p.val)
		}
		p = p.next()
	}
}

func (p Pos) next() Pos {
	nxt := Pos{p.r, p.c, p.val}
	if p.r == 1 {
		nxt.r = p.c + 1
		nxt.c = 1
	} else {
		nxt.r = p.r - 1
		nxt.c = p.c + 1
	}
	nxt.val = (p.val * 252533) % 33554393
	return nxt
}
