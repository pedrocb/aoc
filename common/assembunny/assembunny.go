package assembunny

import (
	"strconv"
)

type Instruction struct {
	A, B string
	Name string
}

func (i Instruction) exec(registers map[string]int, instructions *[]Instruction, currentPosition *int) (int, bool) {
	switch i.Name {
	case "cpy":
		// IF register does not exist, it must be an int
		lvalue, exists := registers[i.A]
		if !exists {
			lvalue, _ = strconv.Atoi(i.A)
		}
		if _, exists := registers[i.B]; exists {
			// Only if b is a register
			registers[i.B] = lvalue
		}
	case "jnz":
		// IF register does not exist, it must be an int
		lvalue, exists := registers[i.A]
		if !exists {
			lvalue, _ = strconv.Atoi(i.A)
		}
		if lvalue != 0 {
			jump, exists := registers[i.B]
			if !exists {
				jump, _ = strconv.Atoi(i.B)
			}
			// - 1 Because loop increases position by 1
			*currentPosition += (jump - 1)
		}
	case "inc":
		registers[i.A] += 1
	case "dec":
		registers[i.A] -= 1
	case "tgl":
		jump, _ := registers[i.A]
		togglePosition := *currentPosition + jump
		if togglePosition < 0 || togglePosition >= len(*instructions) {
			return 0, false
		}
		instructionToToggle := (*instructions)[togglePosition]

		if instructionToToggle.B == "" {
			// One arg instruction
			if instructionToToggle.Name == "inc" {
				(*instructions)[togglePosition].Name = "dec"
			} else {
				(*instructions)[togglePosition].Name = "inc"
			}
		} else {
			if instructionToToggle.Name == "jnz" {
				(*instructions)[togglePosition].Name = "cpy"
			} else {
				(*instructions)[togglePosition].Name = "jnz"
			}
		}
	case "out":
		lvalue, exists := registers[i.A]
		if !exists {
			lvalue, _ = strconv.Atoi(i.A)
		}
		return lvalue, true
	}
	return 0, false
}

func FindOutputLoop(inputInstructions []Instruction, initialState map[string]int, confidenceSize int, loopInterval []int) bool {
	registers := initialState
	loopSize := len(loopInterval)
	numSignals := 0

	// Since tgl changes instructions, better to create a copy
	instructions := make([]Instruction, len(inputInstructions))
	copy(instructions, inputInstructions)

	for currentPosition := 0; currentPosition < len(instructions); currentPosition++ {
		currentInstruction := instructions[currentPosition]
		signal, hasOutput := currentInstruction.exec(registers, &instructions, &currentPosition)
		if !hasOutput {
			continue
		}
		if loopInterval[numSignals%loopSize] == signal {
			numSignals++
			if numSignals/loopSize >= confidenceSize {
				return true
			}
		} else {
			return false
		}

	}
	return false

}

func ExecInstructions(inputInstructions []Instruction, initialState map[string]int) map[string]int {
	registers := initialState

	// Since tgl changes instructions, better to create a copy
	instructions := make([]Instruction, len(inputInstructions))
	copy(instructions, inputInstructions)

	for currentPosition := 0; currentPosition < len(instructions); currentPosition++ {
		currentInstruction := instructions[currentPosition]
		currentInstruction.exec(registers, &instructions, &currentPosition)
	}
	return registers
}
