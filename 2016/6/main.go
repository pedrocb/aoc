package main

import (
	"fmt"
	"strings"
)

func CalcErrorCorrectedMsgModified(originalMessages []string) string {
	messageLength := len(originalMessages[0])

	var correctedMessage strings.Builder

	for charIndex := 0; charIndex < messageLength; charIndex++ {
		charHistogram := make(map[byte]int)
		for _, ogMsg := range originalMessages {
			charHistogram[ogMsg[charIndex]] += 1
		}
		var leastCommonChar byte
		minimumFrequency := 0
		for charCandidate, charFrequency := range charHistogram {
			if charFrequency < minimumFrequency || minimumFrequency == 0 {
				minimumFrequency = charFrequency
				leastCommonChar = charCandidate
			}
		}
		correctedMessage.WriteByte(leastCommonChar)
	}
	return correctedMessage.String()
}

func CalcErrorCorrectedMsg(originalMessages []string) string {
	messageLength := len(originalMessages[0])

	var correctedMessage strings.Builder

	for charIndex := 0; charIndex < messageLength; charIndex++ {
		charHistogram := make(map[byte]int)
		var mostCommonChar byte
		for _, ogMsg := range originalMessages {
			charHistogram[ogMsg[charIndex]] += 1
			if charHistogram[ogMsg[charIndex]] > charHistogram[mostCommonChar] {
				mostCommonChar = ogMsg[charIndex]
			}
		}
		correctedMessage.WriteByte(mostCommonChar)
	}
	return correctedMessage.String()
}

func main() {
	var inputList []string
	for {
		var inputString string
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}

		inputList = append(inputList, inputString)
	}

	// Puzzle 1
	result1 := CalcErrorCorrectedMsg(inputList)
	fmt.Println(result1)

	// Puzzle 2
	result2 := CalcErrorCorrectedMsgModified(inputList)
	fmt.Println(result2)
}
