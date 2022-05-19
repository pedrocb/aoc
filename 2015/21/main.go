package main

import (
	"fmt"
	"math"
)

type character struct {
	hitPoints int
	damage    int
	armor     int
}

type equipment struct {
	cost   int
	damage int
	armor  int
}

var weapons = []equipment{
	equipment{8, 4, 0},
	equipment{10, 5, 0},
	equipment{25, 6, 0},
	equipment{40, 7, 0},
	equipment{74, 8, 0},
}

var armor = []equipment{
	equipment{13, 0, 1},
	equipment{31, 0, 2},
	equipment{53, 0, 3},
	equipment{75, 0, 4},
	equipment{102, 0, 5},
	equipment{0, 0, 0},
}

var rings = []equipment{
	equipment{25, 1, 0},
	equipment{50, 2, 0},
	equipment{100, 3, 0},
	equipment{20, 0, 1},
	equipment{40, 0, 2},
	equipment{80, 0, 3},
	equipment{0, 0, 0},
	equipment{0, 0, 0},
}

func partOne(enemy character) int {
	minimumCost := -1
	for _, candidateWeapon := range weapons {
		cost := candidateWeapon.cost
		for _, candidateArmor := range armor {
			cost += candidateArmor.cost
			if cost < minimumCost || minimumCost == -1 {
				for _, candidateLeftRing := range rings {
					cost += candidateLeftRing.cost
					if cost < minimumCost || minimumCost == -1 {
						for _, candidateRightRing := range rings {
							cost += candidateRightRing.cost
							if cost < minimumCost || minimumCost == -1 {
								myArmor := candidateWeapon.armor + candidateArmor.armor + candidateLeftRing.armor + candidateRightRing.armor
								myDamage := candidateWeapon.damage + candidateArmor.damage + candidateLeftRing.damage + candidateRightRing.damage
								myHp := 100

								myDamagePerTurn := myDamage - enemy.armor
								if myDamagePerTurn < 1 {
									myDamagePerTurn = 1
								}

								myLossPerTurn := enemy.damage - myArmor
								if myLossPerTurn < 1 {
									myLossPerTurn = 1
								}

								turnsToWin := math.Ceil(float64(enemy.hitPoints) / float64(myDamagePerTurn))
								turnsToLose := math.Ceil(float64(myHp) / float64(myLossPerTurn))
								if turnsToLose >= turnsToWin {
									minimumCost = cost
								}
							}
							cost -= candidateRightRing.cost
						}
					}
					cost -= candidateLeftRing.cost
				}
			}
			cost -= candidateArmor.cost
		}
	}
	return minimumCost
}

func partTwo(enemy character) int {
	maximumCost := 0
	for _, candidateWeapon := range weapons {
		for _, candidateArmor := range armor {
			for _, candidateLeftRing := range rings {
				for _, candidateRightRing := range rings {
					cost := candidateWeapon.cost + candidateArmor.cost + candidateLeftRing.cost + candidateRightRing.cost
					if cost > maximumCost {
						myArmor := candidateWeapon.armor + candidateArmor.armor + candidateLeftRing.armor + candidateRightRing.armor
						myDamage := candidateWeapon.damage + candidateArmor.damage + candidateLeftRing.damage + candidateRightRing.damage
						myHp := 100

						myDamagePerTurn := myDamage - enemy.armor
						if myDamagePerTurn < 1 {
							myDamagePerTurn = 1
						}

						myLossPerTurn := enemy.damage - myArmor
						if myLossPerTurn < 1 {
							myLossPerTurn = 1
						}

						turnsToWin := math.Ceil(float64(enemy.hitPoints) / float64(myDamagePerTurn))
						turnsToLose := math.Ceil(float64(myHp) / float64(myLossPerTurn))
						if turnsToLose < turnsToWin {
							maximumCost = cost
						}
					}
				}
			}
		}
	}
	return maximumCost
}

func main() {
	var enemy character
	fmt.Scanf("Hit Points: %d", &enemy.hitPoints)
	fmt.Scanf("Damage: %d", &enemy.damage)
	fmt.Scanf("Armor: %d", &enemy.armor)
	fmt.Println(partOne(enemy))
	fmt.Println(partTwo(enemy))

}
