package main

import (
	"fmt"
)

func partOne(inputList []string) int {
	totalDifference := 0
	for _, currentString := range inputList {
		characterCount := 0
		lenString := len(currentString)
		for indexChar := 1; indexChar < lenString; indexChar += 1 {
			if indexChar == lenString-1 {
				continue
			}

			currentChar := currentString[indexChar]
			if currentChar == '\\' {
				if currentString[indexChar+1] == '"' || currentString[indexChar+1] == '\\' {
					indexChar += 1
				} else if currentString[indexChar+1] == 'x' {
					indexChar += 3
				}
			}
			characterCount += 1
		}
		totalDifference += (lenString - characterCount)
	}
	return totalDifference
}

func partTwo(inputList []string) int {
	totalDifference := 0
	for _, currentString := range inputList {
		lenString := len(currentString)
		totalEncodedChars := 2
		for indexChar := 0; indexChar < lenString; indexChar += 1 {
			currentChar := currentString[indexChar]
			if currentChar == '\\' || currentChar == '"' {
				totalEncodedChars += 1
			}
			totalEncodedChars += 1
		}
		totalDifference += (totalEncodedChars - lenString)
	}
	return totalDifference
}

func main() {
	inputList := []string{}
	var inputString string
	for {
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
		inputList = append(inputList, inputString)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))
}
