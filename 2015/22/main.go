package main

import (
	"fmt"
)

const (
	Easy = 1
	Hard = 2
)
const (
	MagicMissile = 1
	Drain        = 2
	Shield       = 3
	Poison       = 4
	Recharge     = 5
)

var effectsDuraction = map[int]int{
	Shield:   6,
	Poison:   6,
	Recharge: 5,
}

var spells = map[int]int{
	MagicMissile: 53,
	Drain:        73,
	Shield:       113,
	Poison:       173,
	Recharge:     229,
}

type character struct {
	hp     int
	damage int
	armor  int
	mana   int
}

func calcLeastAmountMana(currentMana int, currentEffects map[int]int, minimum *int, enemy character, me character, spellsUsed []int, difficulty int) {
	if currentMana >= *minimum && *minimum != -1 {
		return
	}
	if difficulty == Hard {
		me.hp -= 1
		if me.hp <= 0 {
			return
		}
	}
	// Boss turn
	for effect, numTurns := range currentEffects {
		if numTurns == -1 {
			if effect == Shield {
				me.armor = 0
			}
			delete(currentEffects, effect)
			continue
		} else if numTurns > 0 {
			switch effect {
			case Shield:
				me.armor = 7
			case Poison:
				enemy.hp -= 3
			case Recharge:
				me.mana += 101

			}
		}
		currentEffects[effect] -= 1

	}
	if enemy.hp <= 0 {
		*minimum = currentMana
		return
	}
	me.hp -= (enemy.damage - me.armor)

	if me.hp <= 0 {
		return
	}

	if difficulty == Hard {
		me.hp -= 1
		if me.hp <= 0 {
			return
		}
	}
	// My turn
	for effect, numTurns := range currentEffects {
		if numTurns == -1 {
			if effect == Shield {
				me.armor = 0
			}
			delete(currentEffects, effect)
			continue
		} else if numTurns > 0 {
			switch effect {
			case Shield:
				me.armor = 7
			case Poison:
				enemy.hp -= 3
			case Recharge:
				me.mana += 101

			}
		}
		currentEffects[effect] -= 1

	}
	if enemy.hp <= 0 {
		*minimum = currentMana
		return
	}
	for spell, spellCost := range spells {
		// Check if worth
		if currentMana+spellCost > *minimum && *minimum != -1 {
			continue
		}
		// Check if spell is in effect
		if numTurns, exist := currentEffects[spell]; exist && numTurns > 0 {
			continue
		}
		// Check if possible
		if me.mana-spellCost >= 0 {
			currentEffectsCopy := make(map[int]int)
			for effect, numTurns := range currentEffects {
				currentEffectsCopy[effect] = numTurns
			}
			meCopy := me
			meCopy.mana -= spellCost
			enemyCopy := enemy
			switch spell {
			case MagicMissile:
				enemyCopy.hp -= 4
			case Drain:
				meCopy.hp += 2
				enemyCopy.hp -= 2
			default:
				currentEffectsCopy[spell] = effectsDuraction[spell]
			}

			calcLeastAmountMana(currentMana+spellCost, currentEffectsCopy, minimum, enemyCopy, meCopy, append(spellsUsed, spell), difficulty)
		}
	}

}

func partOne(enemy character) int {
	minimum := -1
	me := character{hp: 50, mana: 500}

	for spell, spellCost := range spells {
		currentEffects := make(map[int]int)
		meCopy := me
		meCopy.mana -= spellCost
		enemyCopy := enemy
		switch spell {
		case MagicMissile:
			enemyCopy.hp -= 4
		case Drain:
			meCopy.hp += 2
			enemyCopy.hp -= 2
		default:
			currentEffects[spell] = effectsDuraction[spell]
		}
		calcLeastAmountMana(spellCost, currentEffects, &minimum, enemyCopy, meCopy, []int{spell}, Easy)
	}
	return minimum
}

func partTwo(enemy character) int {
	minimum := -1
	me := character{hp: 49, mana: 500}

	for spell, spellCost := range spells {
		currentEffects := make(map[int]int)
		meCopy := me
		meCopy.mana -= spellCost
		enemyCopy := enemy
		switch spell {
		case MagicMissile:
			enemyCopy.hp -= 4
		case Drain:
			meCopy.hp += 2
			enemyCopy.hp -= 2
		default:
			currentEffects[spell] = effectsDuraction[spell]
		}
		calcLeastAmountMana(spellCost, currentEffects, &minimum, enemyCopy, meCopy, []int{spell}, Hard)
	}
	return minimum
}

func main() {
	var enemy character
	fmt.Scanf("Hit Points: %d", &enemy.hp)
	fmt.Scanf("Damage: %d", &enemy.damage)
	fmt.Println(partOne(enemy))
	fmt.Println(partTwo(enemy))

}
