package main

import (
	"fmt"
)

type instruction struct {
	register             string
	name                 string
	amount               int
	conditionRegister    string
	conditionInstruction string
	conditionThreshold   int
}

func completeInstructions(inputInstructions []instruction) (map[string]int, int) {
	registers := map[string]int{}
	maximum := 0
	for _, currentInstruction := range inputInstructions {
		conditionRegisterValue, exists := registers[currentInstruction.conditionRegister]
		if !exists {
			conditionRegisterValue = 0
		}
		condition := false
		conditionThreshold := currentInstruction.conditionThreshold
		switch currentInstruction.conditionInstruction {
		case "<":
			condition = conditionRegisterValue < conditionThreshold
		case "<=":
			condition = conditionRegisterValue <= conditionThreshold
		case "==":
			condition = conditionRegisterValue == conditionThreshold
		case ">":
			condition = conditionRegisterValue > conditionThreshold
		case ">=":
			condition = conditionRegisterValue >= conditionThreshold
		case "!=":
			condition = conditionRegisterValue != conditionThreshold
		}
		if condition {
			_, exists := registers[currentInstruction.register]
			if !exists {
				registers[currentInstruction.register] = 0
			}
			amount := currentInstruction.amount
			if currentInstruction.name == "dec" {
				amount *= -1
			}
			registers[currentInstruction.register] += amount
			if registers[currentInstruction.register] > maximum {
				maximum = registers[currentInstruction.register]
			}
		}

	}
	return registers, maximum
}

func findMaximumValue(registers map[string]int) int {
	maximum := 0
	for _, value := range registers {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}

func main() {
	instructions := []instruction{}
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%s %s %d if %s %s %d",
			&inputInstruction.register,
			&inputInstruction.name,
			&inputInstruction.amount,
			&inputInstruction.conditionRegister,
			&inputInstruction.conditionInstruction,
			&inputInstruction.conditionThreshold)
		if err != nil {
			break
		}
		instructions = append(instructions, inputInstruction)
	}
	registers, memoryNeeded := completeInstructions(instructions)

	// Part 1
	result1 := findMaximumValue(registers)
	fmt.Println(result1)

	// Part 2
	result2 := memoryNeeded
	fmt.Println(result2)
}
