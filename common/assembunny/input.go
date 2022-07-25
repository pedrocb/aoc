package assembunny

import (
	"fmt"
)

func ScanInstructions() []Instruction {
	instructions := []Instruction{}
	for {
		var inputInstruction Instruction
		_, err := fmt.Scanf("%s %s", &inputInstruction.Name, &inputInstruction.A)
		if err != nil {
			break
		}

		if inputInstruction.Name == "cpy" || inputInstruction.Name == "jnz" {
			fmt.Scanf("%s", &inputInstruction.B)
		}
		instructions = append(instructions, inputInstruction)
	}
	return instructions
}
