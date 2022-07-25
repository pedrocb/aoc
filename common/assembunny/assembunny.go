package assembunny

import (
	"strconv"
)

type Instruction struct {
	A, B string
	Name string
}

func ExecInstructions(inputInstructions []Instruction, initialState map[string]int) map[string]int {
	registers := initialState

	// Since tgl changes instructions, better to create a copy
	instructions := make([]Instruction, len(inputInstructions))
	copy(instructions, inputInstructions)

	for currentPosition := 0; currentPosition < len(instructions); currentPosition++ {
		currentInstruction := instructions[currentPosition]
		switch currentInstruction.Name {
		case "cpy":
			// IF register does not exist, it must be an int
			lvalue, exists := registers[currentInstruction.A]
			if !exists {
				lvalue, _ = strconv.Atoi(currentInstruction.A)
			}
			if _, exists := registers[currentInstruction.B]; exists {
				// Only if b is a register
				registers[currentInstruction.B] = lvalue
			}
		case "jnz":
			// IF register does not exist, it must be an int
			lvalue, exists := registers[currentInstruction.A]
			if !exists {
				lvalue, _ = strconv.Atoi(currentInstruction.A)
			}
			if lvalue != 0 {
				jump, exists := registers[currentInstruction.B]
				if !exists {
					jump, _ = strconv.Atoi(currentInstruction.B)
				}
				// - 1 Because loop increases position by 1
				currentPosition += (jump - 1)
			}
		case "inc":
			registers[currentInstruction.A] += 1
		case "dec":
			registers[currentInstruction.A] -= 1
		case "tgl":
			jump, _ := registers[currentInstruction.A]
			togglePosition := currentPosition + jump
			if togglePosition < 0 || togglePosition >= len(inputInstructions) {
				continue
			}
			instructionToToggle := instructions[togglePosition]

			if instructionToToggle.B == "" {
				// One arg instruction
				if instructionToToggle.Name == "inc" {
					instructions[togglePosition].Name = "dec"
				} else {
					instructions[togglePosition].Name = "inc"
				}
			} else {
				if instructionToToggle.Name == "jnz" {
					instructions[togglePosition].Name = "cpy"
				} else {
					instructions[togglePosition].Name = "jnz"
				}
			}
		}
	}
	return registers
}
