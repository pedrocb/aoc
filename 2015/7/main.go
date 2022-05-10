package main

import (
	"fmt"
	"strconv"
)

const (
	And    = "AND"
	Or     = "OR"
	LShift = "LSHIFT"
	RShift = "RSHIFT"
	Not    = "NOT"
	Assign = "->"
)

type instruction struct {
	gate     string
	leftArg  string
	rightArg string
	output   string
}

func calcInstruction(currentInstruction instruction, wiresMap map[string]instruction, wiresState map[string]int) (result int) {
	var lvalue, rvalue int
	var exists bool
	lvalue, err := strconv.Atoi(currentInstruction.leftArg)
	if err != nil {
		lvalue, exists = wiresState[currentInstruction.leftArg]
		if !exists {
			lvalue = calcInstruction(wiresMap[currentInstruction.leftArg], wiresMap, wiresState)
			wiresState[currentInstruction.leftArg] = lvalue
		}
	}
	if currentInstruction.rightArg != "" {
		rvalue, err = strconv.Atoi(currentInstruction.rightArg)
		if err != nil {
			rvalue, exists = wiresState[currentInstruction.rightArg]
			if !exists {
				rvalue = calcInstruction(wiresMap[currentInstruction.rightArg], wiresMap, wiresState)
				wiresState[currentInstruction.rightArg] = rvalue
			}
		}
	}
	switch currentInstruction.gate {
	case And:
		result = lvalue & rvalue
	case Or:
		result = lvalue | rvalue
	case RShift:
		result = lvalue >> rvalue
	case LShift:
		result = lvalue << rvalue
	case Not:
		result = ^lvalue
	case Assign:
		result = lvalue
	}
	return
}

func partOne(inputInstructions []instruction) int {
	wiresState := make(map[string]int)
	wiresMap := make(map[string]instruction)
	for _, currentInstruction := range inputInstructions {
		wiresMap[currentInstruction.output] = currentInstruction
	}
	result := calcInstruction(wiresMap["a"], wiresMap, wiresState)
	return result
}

func partTwo(inputInstructions []instruction) int {
	wiresState := make(map[string]int)
	wiresMap := make(map[string]instruction)
	for _, currentInstruction := range inputInstructions {
		wiresMap[currentInstruction.output] = currentInstruction
	}
	wiresState["b"] = 956
	result := calcInstruction(wiresMap["a"], wiresMap, wiresState)
	return result
}

func main() {
	var inputString string
	inputInstructions := []instruction{}
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
		if inputString == Not {
			inputInstruction.gate = Not
			_, err = fmt.Scanf("%s -> %s", &inputInstruction.leftArg, &inputInstruction.output)
		} else {
			inputInstruction.leftArg = inputString

			_, err = fmt.Scanf("%s", &inputString)
			inputInstruction.gate = inputString
			if inputInstruction.gate == Assign {
				_, err = fmt.Scanf("%s", &inputInstruction.output)
				inputInstructions = append([]instruction{inputInstruction}, inputInstructions...)
				continue
			} else {

				_, err = fmt.Scanf("%s -> %s", &inputInstruction.rightArg, &inputInstruction.output)
			}
		}
		inputInstructions = append(inputInstructions, inputInstruction)
	}
	fmt.Println(uint16(partOne(inputInstructions)))
	fmt.Println(uint16(partTwo(inputInstructions)))
}
