package main

import (
	"fmt"
)

type instruction struct {
	orientation byte
	nSteps      int
}

type coordinates struct {
	x, y int
}

func calcDistanceFromPos(pos coordinates) int {
	xDistance := pos.x
	if xDistance < 0 {
		xDistance *= -1
	}
	yDistance := pos.y

	if yDistance < 0 {
		yDistance *= -1
	}
	return xDistance + yDistance
}

func nextStep(currentPosition coordinates, currentOrientation coordinates, inputInstruction instruction) (coordinates, coordinates) {

	prevOrientation := currentOrientation
	switch inputInstruction.orientation {
	case 'L':
		currentOrientation.x = -prevOrientation.y
		currentOrientation.y = prevOrientation.x
	case 'R':
		currentOrientation.x = prevOrientation.y
		currentOrientation.y = -prevOrientation.x
	}
	currentPosition.x += inputInstruction.nSteps * currentOrientation.x
	currentPosition.y += inputInstruction.nSteps * currentOrientation.y
	return currentPosition, currentOrientation
}

func CalcDistance(inputInstructions []instruction) int {
	currentOrientation := coordinates{0, 1}
	currentPosition := coordinates{0, 0}
	for _, currentInstruction := range inputInstructions {
		currentPosition, currentOrientation = nextStep(currentPosition, currentOrientation, currentInstruction)
	}
	return calcDistanceFromPos(currentPosition)
}

func CalcDistanceFirstRevisitedPos(inputInstructions []instruction) int {
	visitedSteps := map[coordinates]bool{coordinates{0, 0}: true}
	currentOrientation := coordinates{0, 1}
	currentPosition := coordinates{0, 0}
	for _, currentInstruction := range inputInstructions {
		previousPosition := currentPosition
		currentPosition, currentOrientation = nextStep(currentPosition, currentOrientation, currentInstruction)
		for previousPosition != currentPosition {
			previousPosition.x += currentOrientation.x
			previousPosition.y += currentOrientation.y
			if _, visited := visitedSteps[previousPosition]; visited {
				return calcDistanceFromPos(previousPosition)
			}
			visitedSteps[previousPosition] = true
		}
	}
	return -1
}

func main() {
	var inputList []instruction
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%c%d", &inputInstruction.orientation, &inputInstruction.nSteps)
		if err != nil {
			break
		}

		inputList = append(inputList, inputInstruction)
		fmt.Scanf(",")
	}

	// Puzzle 1
	result1 := CalcDistance(inputList)
	fmt.Println(result1)

	// Puzzle 2
	result2 := CalcDistanceFirstRevisitedPos(inputList)
	fmt.Println(result2)
}
