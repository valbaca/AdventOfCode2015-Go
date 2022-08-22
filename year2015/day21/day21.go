package day21

// TIL: pulled in my first non-stdlib module for permutations. Even then, Go's lack of generics means instead of an
// actual list of permutations, I get indexes, which I then take and use to generate the actual permutations.
// Also used a scanner for the first time to parse the item data
import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"strings"
	"valbaca.com/advent/elf"
)

type BaseCh struct {
	name   string
	fullHp int
}

// Ch Character, as in Player Character or Non-Player Character
type Ch struct {
	base  BaseCh
	hp    int
	dmg   int
	armor int
}

type Item struct {
	name  string
	cost  int
	dmg   int
	armor int
}

func Part1(input string, yourHp, bossHp, bossDmg, bossArmor int) int {
	you := newCh("You", yourHp, 0, 0)
	boss := newCh("Boss", bossHp, bossDmg, bossArmor)

	weapons, armor, ringPairs := buildItems(input)

	minCostFound := elf.MaxInt
	for _, w := range weapons {
		for _, a := range armor {
			for _, ringPair := range ringPairs {
				left, right := ringPair[0], ringPair[1]
				cost := w.cost + a.cost + left.cost + right.cost
				if cost < minCostFound {
					// FIGHT!!!
					you.equip(w, a, left, right)
					if you.beats(boss) {
						minCostFound = cost
					}
					you.reset()
					boss.resetHp()
				}
			}
		}
	}
	return minCostFound
}

func Part2(input string, yourHp, bossHp, bossDmg, bossArmor int) int {
	you := newCh("You", yourHp, 0, 0)
	boss := newCh("Boss", bossHp, bossDmg, bossArmor)

	weapons, armor, ringPairs := buildItems(input)

	maxCostFound := elf.MinInt
	for _, w := range weapons {
		for _, a := range armor {
			for _, ringPair := range ringPairs {
				left, right := ringPair[0], ringPair[1]
				cost := w.cost + a.cost + left.cost + right.cost
				if cost > maxCostFound {
					// FIGHT!!!
					you.equip(w, a, left, right)
					if !you.beats(boss) {
						maxCostFound = cost
					}
					you.reset()
					boss.resetHp()
				}
			}
		}
	}
	return maxCostFound
}

func buildItems(input string) (weapons []Item, armor []Item, ringPairs [][]Item) {
	// Build items
	weapons = make([]Item, 0, 6-1)
	lines := strings.Split(input, "\n")
	for _, line := range lines[1:6] {
		weapons = append(weapons, itemFromLine(line))
	}
	armor = make([]Item, 0, 13-8)
	for _, line := range lines[8:13] {
		armor = append(armor, itemFromLine(line))
	}
	armor = append(armor, NoItem) // allow no-armor
	rings := make([]Item, 0, 21-15)
	for _, line := range lines[15:21] {
		rings = append(rings, itemFromLine(line))
	}
	rings = append(rings, NoItem) // two hands, two rings
	rings = append(rings, NoItem)
	ringPairs = ringPermutations(rings)
	return weapons, armor, ringPairs
}

func ringPermutations(rings []Item) [][]Item {
	n := len(rings)
	k := 2
	// https://pkg.go.dev/gonum.org/v1/gonum/stat/combin#Permutations
	permIndexes := combin.Permutations(n, k)
	perms := make([][]Item, 0, len(permIndexes))
	for _, v := range permIndexes {
		perms = append(perms, []Item{rings[v[0]], rings[v[1]]})
	}
	return perms
}

func newCh(name string, hp int, dmg int, armor int) *Ch {
	return &Ch{
		base: BaseCh{
			name:   name,
			fullHp: hp,
		},
		hp:    hp,
		dmg:   dmg,
		armor: armor,
	}
}

func (c *Ch) isDead() bool {
	return c.hp <= 0
}

func (c *Ch) beats(other *Ch) bool {
	prevHp := c.hp
	prevOtherHp := other.hp
	var result bool
	for {
		c.hit(other)
		if other.isDead() {
			result = true
			break
		}
		other.hit(c)
		if c.isDead() {
			result = false
			break
		}
	}
	c.hp = prevHp
	other.hp = prevOtherHp
	return result
}

func (c *Ch) hit(other *Ch) {
	dmg := elf.Max(c.dmg-other.armor, 1)
	other.hp -= dmg
}

func (c *Ch) reset() {
	c.resetHp()
	c.dmg = 0
	c.armor = 0
}

func (c *Ch) resetHp() {
	c.hp = c.base.fullHp
}

func (c *Ch) equip(items ...Item) {
	for _, item := range items {
		c.armor += item.armor
		c.dmg += item.dmg
	}
}

func itemFromLine(line string) Item {
	splits := make([]string, 0, 4)
	// https://pkg.go.dev/bufio#example-Scanner-Words
	// https://stackoverflow.com/questions/38026530/a-better-way-to-use-scanner-for-multiple-tokens-per-line
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		splits = append(splits, word)
	}
	if len(splits) == 0 || len(splits) > 4 {
		panic(fmt.Sprintf("error processing item line %v", line))
	}
	return Item{
		name:  splits[0],
		cost:  elf.UnsafeAtoi(splits[1]),
		dmg:   elf.UnsafeAtoi(splits[2]),
		armor: elf.UnsafeAtoi(splits[3]),
	}
}

var NoItem Item = Item{"NO ITEM", 0, 0, 0}
