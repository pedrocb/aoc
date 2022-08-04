package main

import (
	"fmt"
)

func CalcScore(inputString string) int {
	score := 0
	currentLevel := 0
	inGarbage := false
	for index := 0; index < len(inputString); index++ {
		switch inputString[index] {
		case '{':
			if !inGarbage {
				currentLevel++
			}
		case '!':
			index++
		case '<':
			if !inGarbage {
				inGarbage = true
			}
		case '>':
			inGarbage = false
		case '}':
			if !inGarbage {
				score += currentLevel
				currentLevel--
			}
		}
	}
	return score
}

func CalcGarbageChars(inputString string) int {
	counter := 0
	inGarbage := false
	for index := 0; index < len(inputString); index++ {
		switch inputString[index] {
		case '!':
			index++
		case '<':
			if !inGarbage {
				inGarbage = true
			} else {
				counter++
			}
		case '>':
			inGarbage = false
		default:
			if inGarbage {
				counter++
			}
		}
	}
	return counter
}

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)

	// Part 1
	result1 := CalcScore(inputString)
	fmt.Println(result1)

	// Part 2
	result2 := CalcGarbageChars(inputString)
	fmt.Println(result2)
}
