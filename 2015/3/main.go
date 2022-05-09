package main

import (
	"fmt"
)

func partOne(inputArray []byte) int {
	arrayLength := len(inputArray)
	var currentPosition = arrayLength / 2 * arrayLength / 2
	grid := make([]bool, arrayLength*arrayLength)
	grid[currentPosition] = true
	counter := 1
	for _, step := range inputArray {
		if step == '>' {
			currentPosition = currentPosition + 1
		} else if step == '<' {
			currentPosition = currentPosition - 1
		} else if step == '^' {
			currentPosition = currentPosition + arrayLength
		} else if step == 'v' {
			currentPosition = currentPosition - arrayLength
		}
		if !grid[currentPosition] {
			grid[currentPosition] = true
			counter += 1
		}
	}
	return counter
}

func partTwo(inputArray []byte) int {
	arrayLength := len(inputArray)
	var robotCurrentPosition = arrayLength * arrayLength
	var santaCurrentPosition = arrayLength * arrayLength
	grid := make([]bool, arrayLength*2*arrayLength*2)
	grid[santaCurrentPosition] = true
	counter := 1
	for index, step := range inputArray {
		var currentPosition *int
		if index%2 == 0 {
			currentPosition = &santaCurrentPosition
		} else {
			currentPosition = &robotCurrentPosition
		}
		if step == '>' {
			*currentPosition = *currentPosition + 1
		} else if step == '<' {
			*currentPosition = *currentPosition - 1
		} else if step == '^' {
			*currentPosition = *currentPosition + (arrayLength * 2)
		} else if step == 'v' {
			*currentPosition = *currentPosition - (arrayLength * 2)
		}
		if !grid[*currentPosition] {
			grid[*currentPosition] = true
			counter += 1
		}
	}
	return counter
}

func main() {
	inputArray := []byte{}
	var inputChar byte
	for {
		_, err := fmt.Scanf("%c", &inputChar)
		if err != nil {
			break
		}
		inputArray = append(inputArray, inputChar)
	}
	fmt.Println(partOne(inputArray))
	fmt.Println(partTwo(inputArray))
}
