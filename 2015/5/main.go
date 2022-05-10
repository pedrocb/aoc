package main

import (
	"fmt"
	"strings"
)

func partOne(inputArray []string) int {
	counter := 0
	forbiddenPairs := []string{"ab", "cd", "pq", "xy"}
	for _, candidate := range inputArray {
		vowelCount := 0
		hasDoubleLetter := false
		hasForbiddenPairs := false
		for index, currentChar := range candidate {
			if strings.ContainsRune("aeiou", currentChar) {
				vowelCount += 1
			}
			if index > 0 {
				if !hasDoubleLetter && byte(currentChar) == candidate[index-1] {
					hasDoubleLetter = true
				}
				for _, forbiddenPair := range forbiddenPairs {
					if candidate[index-1:index+1] == forbiddenPair {
						hasForbiddenPairs = true
						break
					}
				}
				if hasForbiddenPairs {
					break
				}
			}
		}
		if vowelCount >= 3 && hasDoubleLetter && !hasForbiddenPairs {
			counter += 1
		}
	}
	return counter
}

func partTwo(inputArray []string) int {
	counter := 0
	for _, candidate := range inputArray {
		hasRuleOne := false
		hasRuleTwo := false
		lenCandidate := len(candidate)
		pairsMap := make(map[string]bool)
		var lastPair string
		for index, _ := range candidate {
			if index == 0 {
				continue
			}
			currentPair := candidate[index-1 : index+1]
			if !hasRuleOne && pairsMap[currentPair] && currentPair != lastPair {
				hasRuleOne = true
			}
			pairsMap[currentPair] = true
			lastPair = currentPair

			if !hasRuleTwo && index > 0 && index < lenCandidate-1 {
				if candidate[index-1] == candidate[index+1] {
					hasRuleTwo = true
				}
			}
			if hasRuleOne && hasRuleTwo {
				counter += 1
				break
			}
		}
	}
	return counter
}

func main() {
	inputArray := []string{}
	var inputString string
	for {
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
		inputArray = append(inputArray, inputString)
	}
	fmt.Println(partOne(inputArray))
	fmt.Println(partTwo(inputArray))
}
