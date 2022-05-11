package main

import (
	"fmt"
	"strconv"
)

func lookAndSay(input []byte) []byte {
	var output []byte
	charSequence := input[0]
	sequenceSize := 0
	for _, currentChar := range input {
		if currentChar != charSequence {
			output = append(output, strconv.Itoa(sequenceSize)[0])
			output = append(output, charSequence)
			charSequence = currentChar
			sequenceSize = 1
		} else {
			sequenceSize += 1
		}
	}
	output = append(output, strconv.Itoa(sequenceSize)[0])
	output = append(output, charSequence)
	return output
}

func partOne(inputString string) int {
	input := []byte(inputString)
	var output []byte
	for i := 0; i < 40; i++ {
		output = lookAndSay(input)
		input = output
	}
	return len(output)
}

func partTwo(inputString string) int {
	input := []byte(inputString)
	var output []byte
	for i := 0; i < 50; i++ {
		output = lookAndSay(input)
		input = output
	}
	return len(output)
}

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)
	fmt.Println(partOne(inputString))
	fmt.Println(partTwo(inputString))
}
