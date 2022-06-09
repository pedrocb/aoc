package main

import (
	"fmt"
)

func countDigits(number int) int {
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count

}

func DecompressLength(content string, recursively bool) int {
	count := 0
	for currentCharIndex := 0; currentCharIndex < len(content); currentCharIndex++ {
		var numCharacters, repeat int
		_, err := fmt.Sscanf(content[currentCharIndex:], "(%dx%d)", &numCharacters, &repeat)
		if err == nil {
			// Size of marker (x,y) = 3 symbols + number of digits of both numbers
			markerSize := 3 + countDigits(numCharacters) + countDigits(repeat)
			if recursively {
				count += DecompressLength(content[currentCharIndex+markerSize:currentCharIndex+markerSize+numCharacters], true) * repeat
			} else {
				count += numCharacters * repeat
			}
			// -1 because loop increases
			currentCharIndex += numCharacters + markerSize - 1
		} else {
			count += 1
		}
	}

	return count
}

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)

	result1 := DecompressLength(inputString, false)
	fmt.Println(result1)

	result2 := DecompressLength(inputString, true)
	fmt.Println(result2)
}
