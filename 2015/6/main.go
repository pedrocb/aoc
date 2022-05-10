package main

import (
	"fmt"
)

const (
	TurnOff = 0
	TurnOn  = 1
	Toggle  = 2
)

type instruction struct {
	command        int
	startX, startY int
	endX, endY     int
}

func partOne(inputInstructions []instruction) int {
	counter := 0
	grid := make([]bool, 1000*1000)

	for _, currentInstruction := range inputInstructions {
		for currentX := currentInstruction.startX; currentX <= currentInstruction.endX; currentX += 1 {
			for currentY := currentInstruction.startY; currentY <= currentInstruction.endY; currentY += 1 {
				currentPos := currentY*1000 + currentX
				lightState := grid[currentPos]
				if currentInstruction.command == TurnOff {
					grid[currentPos] = false
				} else if currentInstruction.command == TurnOn {
					grid[currentPos] = true
				} else {
					grid[currentPos] = !lightState
				}
				if !lightState && grid[currentPos] {
					counter += 1
				} else if lightState && !grid[currentPos] {
					counter -= 1
				}
			}
		}
	}
	return counter
}

func partTwo(inputInstructions []instruction) int {
	counter := 0
	grid := make([]int, 1000*1000)

	for _, currentInstruction := range inputInstructions {
		for currentX := currentInstruction.startX; currentX <= currentInstruction.endX; currentX += 1 {
			for currentY := currentInstruction.startY; currentY <= currentInstruction.endY; currentY += 1 {
				currentPos := currentY*1000 + currentX
				var delta int
				if currentInstruction.command == TurnOff {
					delta = -1
					if grid[currentPos] == 0 {
						delta = 0
					}
				} else if currentInstruction.command == TurnOn {
					delta = +1
				} else {
					delta = +2
				}
				grid[currentPos] += delta
				counter += delta
			}
		}
	}
	return counter
}

func main() {
	var inputString string
	instructionsList := []instruction{}
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
		if inputString == "toggle" {
			inputInstruction.command = Toggle
		} else {
			_, err = fmt.Scanf("%s", &inputString)
			if inputString == "on" {
				inputInstruction.command = TurnOn
			} else {
				inputInstruction.command = TurnOff
			}
		}
		_, err = fmt.Scanf("%d,%d through %d,%d",
			&inputInstruction.startX, &inputInstruction.startY,
			&inputInstruction.endX, &inputInstruction.endY)
		instructionsList = append(instructionsList, inputInstruction)
	}
	fmt.Println(partOne(instructionsList))
	fmt.Println(partTwo(instructionsList))
}
