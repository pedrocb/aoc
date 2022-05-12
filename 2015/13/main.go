package main

import (
	"fmt"
	"strings"
)

type pair struct {
	first     string
	second    string
	happiness int
}

func calcHappiness(partition []string, happinessPairs map[string]map[string]int) int {
	counter := 0
	for index, current := range partition {
		indexToCompare := index - 1
		if index == 0 {
			indexToCompare = len(partition) - 1
		}
		if current == "me" || partition[indexToCompare] == "me" {
			continue
		}
		counter += happinessPairs[current][partition[indexToCompare]]
	}
	return counter
}

func maxFromPartition(currentPartition []string, filter int, familyMembers []string, n int, happinessPairs map[string]map[string]int) int {
	maximum := 0
	if len(currentPartition) == n {
		return calcHappiness(currentPartition, happinessPairs)
	} else {
		for index, member := range familyMembers {
			// If member not seated yet
			if filter&(1<<index) == 0 {
				candidate := maxFromPartition(append(currentPartition, member), filter|(1<<index), familyMembers, n, happinessPairs)
				if candidate > maximum {
					maximum = candidate
				}
			}
		}
	}

	return maximum
}

func partOne(inputList []pair) int {
	happinessPairs := make(map[string]map[string]int)
	var familyMembers []string
	for _, inputPair := range inputList {
		_, exists := happinessPairs[inputPair.first]
		if !exists {
			happinessPairs[inputPair.first] = make(map[string]int)
			familyMembers = append(familyMembers, inputPair.first)
		}
		happinessPairs[inputPair.first][inputPair.second] += inputPair.happiness

		_, exists = happinessPairs[inputPair.second]
		if !exists {
			happinessPairs[inputPair.second] = make(map[string]int)
			familyMembers = append(familyMembers, inputPair.second)
		}
		happinessPairs[inputPair.second][inputPair.first] += inputPair.happiness
	}
	return maxFromPartition([]string{}, 0, familyMembers, len(familyMembers), happinessPairs)
}

func partTwo(inputList []pair) int {
	happinessPairs := make(map[string]map[string]int)
	var familyMembers []string
	for _, inputPair := range inputList {
		_, exists := happinessPairs[inputPair.first]
		if !exists {
			happinessPairs[inputPair.first] = make(map[string]int)
			familyMembers = append(familyMembers, inputPair.first)

		}
		happinessPairs[inputPair.first][inputPair.second] += inputPair.happiness

		_, exists = happinessPairs[inputPair.second]
		if !exists {
			happinessPairs[inputPair.second] = make(map[string]int)
			familyMembers = append(familyMembers, inputPair.second)
		}
		happinessPairs[inputPair.second][inputPair.first] += inputPair.happiness
	}
	familyMembers = append(familyMembers, "me")
	return maxFromPartition([]string{}, 0, familyMembers, len(familyMembers), happinessPairs)
}

func main() {
	var inputList []pair
	for {
		var inputPair pair
		var verb string
		_, err := fmt.Scanf("%s would %s %d happiness units by sitting next to %s", &inputPair.first, &verb, &inputPair.happiness, &inputPair.second)
		inputPair.second = strings.TrimSuffix(inputPair.second, ".")
		if err != nil {
			break
		}
		if verb == "lose" {
			inputPair.happiness *= -1
		}

		inputList = append(inputList, inputPair)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))
}
