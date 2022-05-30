package main

import (
	"fmt"
)

func convertVerticalToHorizontal(inputSizes [][3]int) [][3]int {
	result := [][3]int{}
	for i := 0; i < len(inputSizes)/3; i++ {
		currentRow := i * 3
		for j := 0; j < 3; j++ {
			triangleCandidate := [3]int{inputSizes[currentRow][j], inputSizes[currentRow+1][j], inputSizes[currentRow+2][j]}
			result = append(result, triangleCandidate)
		}
	}
	return result
}

func CalcNumberPossibleTriangles(inputSizes [][3]int) int {
	counter := 0
	for _, triangleSides := range inputSizes {
		isPossible := true
		for currentSide := 0; currentSide < 3; currentSide++ {
			otherSidesSum := 0
			for otherSide := 0; otherSide < 3; otherSide++ {
				if currentSide == otherSide {
					continue
				}
				otherSidesSum += triangleSides[otherSide]
			}
			if otherSidesSum <= triangleSides[currentSide] {
				isPossible = false
				break
			}
		}
		if isPossible {
			counter += 1
		}
	}
	return counter
}

func main() {
	var inputSizes [][3]int
	for {
		var inputSize [3]int
		_, err := fmt.Scanf("%d %d %d", &inputSize[0], &inputSize[1], &inputSize[2])
		if err != nil {
			break
		}
		inputSizes = append(inputSizes, inputSize)
	}

	// Puzzle 1
	result1 := CalcNumberPossibleTriangles(inputSizes)
	fmt.Println(result1)

	// Puzzle 2
	result2 := CalcNumberPossibleTriangles(convertVerticalToHorizontal(inputSizes))
	fmt.Println(result2)
}
