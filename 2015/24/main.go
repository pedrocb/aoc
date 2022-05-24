package main

import (
	"fmt"
)

func calcMinimumQE(combination []int, currentWeight int, combinationMaxSize int, weightTarget int, weights []int) int {
	if len(combination) == combinationMaxSize && currentWeight == weightTarget {
		mde := 1
		for _, weightIndex := range combination {
			mde *= weights[weightIndex]
		}
		fmt.Println(combination)
		return mde
	} else if currentWeight > weightTarget {
		return -1
	} else if len(combination) < combinationMaxSize {
		minimum := -1
		lastWeight := -1
		if len(combination) > 0 {
			lastWeight = combination[len(combination)-1]
		}
		for indexWeight := lastWeight + 1; indexWeight < len(weights); indexWeight++ {
			weight := weights[indexWeight]
			if currentWeight+weight <= weightTarget {
				candidate := calcMinimumQE(append(combination, indexWeight), currentWeight+weight, combinationMaxSize, weightTarget, weights)
				if candidate != -1 && (minimum == -1 || candidate < minimum) {
					minimum = candidate
				}
			}
		}
		return minimum
	}
	return -1
}

func partOne(inputList []int) int {
	sumWeights := 0
	for _, weight := range inputList {
		sumWeights += weight
	}
	sumWeights /= 3

	partitionSize := 2
	for {
		candidateMinimum := calcMinimumQE([]int{}, 0, partitionSize, sumWeights, inputList)
		if candidateMinimum != -1 {
			return candidateMinimum
		}
		partitionSize += 1
	}
	return -1
}

func partTwo(inputList []int) int {
	sumWeights := 0
	for _, weight := range inputList {
		sumWeights += weight
	}
	sumWeights /= 4

	partitionSize := 2
	for {
		candidateMinimum := calcMinimumQE([]int{}, 0, partitionSize, sumWeights, inputList)
		if candidateMinimum != -1 {
			return candidateMinimum
		}
		partitionSize += 1
	}
	return -1
}

func main() {
	var inputList []int
	for {
		var inputNumber int
		_, err := fmt.Scanf("%d", &inputNumber)
		if err != nil {
			break
		}
		inputList = append(inputList, inputNumber)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))

}
