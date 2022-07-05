package main

import (
	"fmt"
	"strconv"
)

type instruction struct {
	a, b string
	name string
}

func ExecInstructions(inputInstructions []instruction, initialState map[string]int) map[string]int {
	registers := initialState
	for currentPosition := 0; currentPosition < len(inputInstructions); currentPosition++ {
		currentInstruction := inputInstructions[currentPosition]
		switch currentInstruction.name {
		case "cpy":
			// IF register does not exist, it must be an int
			lvalue, exists := registers[currentInstruction.a]
			if !exists {
				lvalue, _ = strconv.Atoi(currentInstruction.a)
			}
			registers[currentInstruction.b] = lvalue
		case "jnz":
			// IF register does not exist, it must be an int
			lvalue, exists := registers[currentInstruction.a]
			if !exists {
				lvalue, _ = strconv.Atoi(currentInstruction.a)
			}
			if lvalue != 0 {
				jump, _ := strconv.Atoi(currentInstruction.b)
				// - 1 Because loop increases position by 1
				currentPosition += (jump - 1)
			}
		case "inc":
			registers[currentInstruction.a] += 1
		case "dec":
			registers[currentInstruction.a] -= 1
		}
	}
	return registers
}

func main() {
	inputInstructions := []instruction{}
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%s %s", &inputInstruction.name, &inputInstruction.a)
		if err != nil {
			break
		}

		if inputInstruction.name == "cpy" || inputInstruction.name == "jnz" {
			fmt.Scanf("%s", &inputInstruction.b)
		}
		inputInstructions = append(inputInstructions, inputInstruction)
	}

	// Part One
	result1 := ExecInstructions(inputInstructions, map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	})["a"]
	fmt.Println(result1)

	// Part Two
	result2 := ExecInstructions(inputInstructions, map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	})["a"]
	fmt.Println(result2)
}
