package main

import (
	"fmt"
)

func partOne(inputArray []byte) int {
	counter := 0
	for _, inputChar := range inputArray {
		if inputChar == '(' {
			counter += 1
		} else if inputChar == ')' {
			counter -= 1
		}
	}
	return counter
}

func partTwo(inputArray []byte) int {
	counter := 0
	for index, inputChar := range inputArray {
		if inputChar == '(' {
			counter += 1
		} else if inputChar == ')' {
			counter -= 1
		}
		if counter == -1 {
			return index + 1
		}
	}
	return -1
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
