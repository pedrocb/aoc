package main

import (
	"fmt"
)

func validatePassword(inputChars []int) bool {
	forbiddenLetters := []int{int('i'), int('o'), int('l')}
	containsStraight := false
	pairsCount := 0
	lenInput := len(inputChars)
	for index, currentChar := range inputChars {
		for _, forbiddenLetter := range forbiddenLetters {
			if currentChar == forbiddenLetter {
				return false
			}
		}
		if !containsStraight &&
			index > 1 &&
			inputChars[index] == inputChars[index-1]+1 &&
			inputChars[index-1] == inputChars[index-2]+1 {
			containsStraight = true
		} else if pairsCount < 2 &&
			index > 0 &&
			inputChars[index] == inputChars[index-1] &&
			(index == lenInput-1 ||
				inputChars[index] != inputChars[index+1]) {
			pairsCount += 1
		}

	}

	return (pairsCount > 1 && containsStraight)

}

func newCandidate(inputChars []int) []int {
	inputSize := len(inputChars)
	for index := inputSize - 1; index >= 0; index-- {
		inputChars[index] += 1
		// if char overflows 'z'
		if inputChars[index] == int('z')+1 {
			inputChars[index] = int('a')
		} else {
			return inputChars
		}
	}
	return inputChars
}

func partOne(inputString string) string {
	var inputChars []int
	for _, char := range inputString {
		inputChars = append(inputChars, int(char))
	}
	candidate := newCandidate(inputChars)
	for !validatePassword(candidate) {
		candidate = newCandidate(candidate)
	}
	var result string
	for _, char := range candidate {
		result += string(char)
	}
	return result
}

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)
	resultOne := partOne(inputString)
	fmt.Println(resultOne)
	fmt.Println(partOne(resultOne))
}
