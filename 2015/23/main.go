package main

import (
	"fmt"
	"strings"
)

type instruction struct {
	instructionType string
	register        string
	offset          int
}

func partOne(inputList []instruction) int {
	currentInstructionIndex := 0
	nInstructions := len(inputList)
	state := map[string]int{"a": 0, "b": 0}
	for currentInstructionIndex < nInstructions {
		currentInstruction := inputList[currentInstructionIndex]
		switch currentInstruction.instructionType {
		case "hlf":
			state[currentInstruction.register] /= 2
		case "tpl":
			state[currentInstruction.register] *= 3
		case "inc":
			state[currentInstruction.register] += 1
		case "jmp":
			currentInstructionIndex += currentInstruction.offset
			continue
		case "jie":
			if state[currentInstruction.register]%2 == 0 {
				currentInstructionIndex += currentInstruction.offset
				continue
			}
		case "jio":
			if state[currentInstruction.register] == 1 {
				currentInstructionIndex += currentInstruction.offset
				continue
			}
		}
		currentInstructionIndex += 1
	}
	return state["b"]
}

func partTwo(inputList []instruction) int {
	currentInstructionIndex := 0
	nInstructions := len(inputList)
	state := map[string]int{"a": 1, "b": 0}
	for currentInstructionIndex < nInstructions {
		currentInstruction := inputList[currentInstructionIndex]
		switch currentInstruction.instructionType {
		case "hlf":
			state[currentInstruction.register] /= 2
		case "tpl":
			state[currentInstruction.register] *= 3
		case "inc":
			state[currentInstruction.register] += 1
		case "jmp":
			currentInstructionIndex += currentInstruction.offset
			continue
		case "jie":
			if state[currentInstruction.register]%2 == 0 {
				currentInstructionIndex += currentInstruction.offset
				continue
			}
		case "jio":
			if state[currentInstruction.register] == 1 {
				currentInstructionIndex += currentInstruction.offset
				continue
			}
		}
		currentInstructionIndex += 1
	}
	return state["b"]
}

func main() {
	var inputList []instruction
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%s", &inputInstruction.instructionType)
		if err != nil {
			break
		}
		if inputInstruction.instructionType == "jmp" {
			fmt.Scanf("%d", &inputInstruction.offset)
		} else {
			fmt.Scanf("%s", &inputInstruction.register)
			if inputInstruction.instructionType == "jio" || inputInstruction.instructionType == "jie" {
				inputInstruction.register = strings.TrimRight(inputInstruction.register, ",")
				fmt.Scanf("%d", &inputInstruction.offset)
			}
		}
		inputList = append(inputList, inputInstruction)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))

}
