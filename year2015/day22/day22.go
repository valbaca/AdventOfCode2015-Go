package day22

/*
TIL: spent WAY too long debugging due to a single `:=` within a loop that should've been a `=`
In hindsight, it's super obvious that I was creating new vars that were shadowing the outer scope, but I'm not sure
if I like how implicit it was. Didn't think I'd ever miss Python's `global` and `nonlocal`
There's a reason why JS uses var/let/const and Kotlin uses var/val.

I also thought using pointers for the spell array would give better perf, but it ended up actually taking longer!
*/
import (
	"strconv"
	"valbaca.com/advent/elf"
)

func Part1() string {
	return strconv.Itoa(run([]Spell{}, elf.MaxInt, false))
}

func Part2() string {
	return strconv.Itoa(run([]Spell{}, elf.MaxInt, true))
}

func run(spells []Spell, minFound int, hardMode bool) int {
	result, cost := battle(spells, minFound, hardMode)
	if result == failure || (result == victory && cost >= minFound) {
		return minFound // terminate: died or did worse than best
	} else if result == victory && cost < minFound {
		println("New Min found!", cost)
		return cost
	}
	for _, spell := range Spells {
		subMinFound := run(append(spells, spell), minFound, hardMode)
		if subMinFound < minFound {
			minFound = subMinFound
		}
	}
	return minFound
}

type BattleResult int

const (
	failure   BattleResult = iota // Died, Invalid (went below 0 mana) or went over minFound
	stalemate                     // undetermined victor, continue with search
	victory                       // Compare cost
)

type Effect struct {
	spell    Spell
	duration int
}

func battle(spells []Spell, minFound int, hardMode bool) (result BattleResult, cost int) {
	you := You{50, 500, 0}
	boss := Boss{58, 9}
	var effects []Effect
	spent := 0
	for _, spell := range spells {
		if hardMode {
			you.getHit(1)
			if you.isDead() {
				return failure, spent
			}
		}
		// Each loop is one round: effects, you cast, effects again, boss hits back
		you, boss, effects = apply(you, boss, effects)
		if boss.isDead() {
			return victory, spent
		}
		spent += spell.cost
		if spent > minFound {
			return failure, spent
		}
		you.mana -= spell.cost
		spellInActiveSet := false
		for _, effect := range effects {
			if spell == effect.spell {
				spellInActiveSet = true
			}
		}
		if spellInActiveSet || you.mana < 0 {
			return failure, spent
		}

		// cast the spell!
		switch spell {
		case MagicMissile:
			boss.hp -= 4
		case Drain:
			you.hp += 2
			boss.hp -= 2
		case Shield:
			effects = append(effects, Effect{spell, 6})
		case Poison:
			effects = append(effects, Effect{spell, 6})
		case Recharge:
			effects = append(effects, Effect{spell, 5})
		}

		// end of your turn
		if boss.isDead() {
			return victory, spent
		}

		// Boss' turn
		you, boss, effects = apply(you, boss, effects)
		if boss.isDead() {
			return victory, spent
		}
		you.getHit(boss.dmg)
		if you.isDead() {
			return failure, spent
		}
	}
	return stalemate, spent
}

func apply(you You, boss Boss, effects []Effect) (You, Boss, []Effect) {
	var nextEffects []Effect
	for _, e := range effects {
		switch e.spell {
		case Shield:
			you.armor = 7
		case Poison:
			boss.hp -= 3
		case Recharge:
			you.mana += 101
		}
		if e.duration > 1 {
			nextEffects = append(nextEffects, Effect{e.spell, e.duration - 1})
		} else if e.spell == Shield && e.duration == 1 {
			you.armor = 0
		}
	}
	return you, boss, nextEffects
}

type Spell struct {
	name string
	cost int
}

var (
	MagicMissile = Spell{"MagicMissile", 53}
	Drain        = Spell{"Drain", 73}
	Shield       = Spell{"Shield", 113}
	Poison       = Spell{"Poison", 173}
	Recharge     = Spell{"Recharge", 229}
	Spells       = []Spell{MagicMissile, Drain, Shield, Poison, Recharge}
)

type You struct {
	hp    int
	mana  int
	armor int
}

func (y *You) getHit(dmg int) {
	y.hp -= elf.Max(dmg-y.armor, 1)
}

func (y *You) isDead() bool {
	return y.hp <= 0
}

type Boss struct {
	hp  int
	dmg int
}

func (b *Boss) isDead() bool {
	return b.hp <= 0
}
