package main

import (
	"fmt"
	"math"
)

func CalcCode(inputInstructions []string, keyPad []byte, initialKey int) string {
	currentKey := initialKey
	keyPadSize := int(math.Sqrt(float64(len(keyPad))))
	currentX := currentKey % keyPadSize
	currentY := currentKey / keyPadSize
	resultCode := []byte{}
	deltaX := 0
	deltaY := 0

	for _, instruction := range inputInstructions {
		for _, command := range instruction {
			switch command {
			case 'U':
				deltaX = 0
				deltaY = -1
			case 'L':
				deltaX = -1
				deltaY = 0
			case 'R':
				deltaX = 1
				deltaY = 0
			case 'D':
				deltaX = 0
				deltaY = 1
			}

			newX := currentX + deltaX
			newY := currentY + deltaY

			if newX < 0 || newX >= keyPadSize || newY < 0 || newY >= keyPadSize {
				continue
			} else if keyPad[newY*keyPadSize+newX] == '#' {
				continue
			} else {
				currentX = newX
				currentY = newY
			}
		}
		button := keyPad[currentY*keyPadSize+currentX]
		resultCode = append(resultCode, button)
	}

	return string(resultCode)
}

func main() {
	var inputInstructions []string
	for {
		var inputInstruction string
		_, err := fmt.Scanf("%s", &inputInstruction)
		if err != nil {
			break
		}
		inputInstructions = append(inputInstructions, inputInstruction)
	}

	// Puzzle 1
	result1 := CalcCode(inputInstructions, []byte{
		'1', '2', '3',
		'4', '5', '6',
		'7', '8', '9',
	}, 4)
	fmt.Println(result1)

	// Puzzle 2
	result2 := CalcCode(inputInstructions, []byte{
		'#', '#', '1', '#', '#',
		'#', '2', '3', '4', '#',
		'5', '6', '7', '8', '9',
		'#', 'A', 'B', 'C', '#',
		'#', '#', 'D', '#', '#',
	}, 10)
	fmt.Println(result2)
}
