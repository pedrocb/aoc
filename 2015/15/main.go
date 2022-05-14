package main

import (
	"fmt"
)

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
	name       string
}

func maxScoreDistribution(currentN int, currentDistribution []int, inputList []ingredient, leftToDistribute int, limitCalories int) int {
	if currentN == len(inputList)-1 {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0
		for distributionIndex, distribution := range currentDistribution {
			capacity += distribution * inputList[distributionIndex].capacity
			durability += distribution * inputList[distributionIndex].durability
			flavor += distribution * inputList[distributionIndex].flavor
			texture += distribution * inputList[distributionIndex].texture
			calories += distribution * inputList[distributionIndex].calories
		}
		capacity += leftToDistribute * inputList[len(inputList)-1].capacity
		durability += leftToDistribute * inputList[len(inputList)-1].durability
		flavor += leftToDistribute * inputList[len(inputList)-1].flavor
		texture += leftToDistribute * inputList[len(inputList)-1].texture
		calories += leftToDistribute * inputList[len(inputList)-1].calories

		if limitCalories != -1 && calories != limitCalories {
			return 0
		}
		if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
			return 0
		}
		return capacity * durability * flavor * texture
	} else {
		maximum := 0
		for i := 0; i < leftToDistribute; i++ {
			newDistribution := make([]int, len(inputList))
			copy(newDistribution, currentDistribution)
			newDistribution[currentN] = i
			candidate := maxScoreDistribution(currentN+1, newDistribution, inputList, leftToDistribute-i, limitCalories)
			if candidate > maximum {
				maximum = candidate
			}
		}
		return maximum
	}
}

func partOne(inputList []ingredient) int {
	maximum := maxScoreDistribution(0, make([]int, len(inputList)), inputList, 100, -1)
	return maximum
}

func partTwo(inputList []ingredient) int {
	maximum := maxScoreDistribution(0, make([]int, len(inputList)), inputList, 100, 500)
	return maximum
}

func main() {
	var inputList []ingredient
	for {
		var inputIngredient ingredient
		_, err := fmt.Scanf("%s capacity %d, durability %d, flavor %d, texture %d, calories %d\n",
			&inputIngredient.name, &inputIngredient.capacity, &inputIngredient.durability, &inputIngredient.flavor, &inputIngredient.texture, &inputIngredient.calories)
		if err != nil {
			break

		}

		inputList = append(inputList, inputIngredient)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))

}
