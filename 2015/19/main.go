package main

import (
	"fmt"
)

type replacement struct {
	from, to string
}

type node struct {
	data          string
	numExpansions int
	childs        []node
}

var maxLenghFromAtom = 2

func partOne(inputList []replacement, inputString string) int {
	replacementMap := map[string][]string{}
	for _, inputReplacement := range inputList {
		_, exists := replacementMap[inputReplacement.from]
		if !exists {
			replacementMap[inputReplacement.from] = []string{}
		}
		replacementMap[inputReplacement.from] = append(replacementMap[inputReplacement.from], inputReplacement.to)
	}
	cache := map[string]bool{}
	result := 0

	for currentChar := 0; currentChar < len(inputString); currentChar++ {
		for possibleLength := 0; possibleLength < maxLenghFromAtom; possibleLength++ {
			if currentChar+possibleLength >= len(inputString) {
				break
			}
			atomCandidate := inputString[currentChar : currentChar+possibleLength+1]
			replacementTos, replacementExists := replacementMap[atomCandidate]
			if !replacementExists {
				continue
			}
			for _, atomReplacement := range replacementTos {
				candidate := inputString[:currentChar] + atomReplacement + inputString[currentChar+possibleLength+1:len(inputString)]
				if _, exists := cache[candidate]; !exists {
					cache[candidate] = true
					result += 1
				}

			}
		}

	}
	return result
}

func partTwo(inputList []replacement, inputString string) int {
	root := &node{data: inputString}
	nodesToTraverse := []node{*root}
	replacementMap := map[string][]string{}
	cache := map[string]bool{}
	maxLengthToAtom := 0
	for _, inputReplacement := range inputList {
		_, exists := replacementMap[inputReplacement.to]
		if !exists {
			replacementMap[inputReplacement.to] = []string{}
		}
		if len(inputReplacement.to) > maxLengthToAtom {
			maxLengthToAtom = len(inputReplacement.to)
		}
		replacementMap[inputReplacement.to] = append(replacementMap[inputReplacement.to], inputReplacement.from)
	}
	for {
		candidate := nodesToTraverse[0]

		nodesToTraverse = nodesToTraverse[1:]
		if candidate.data == "e" {
			return candidate.numExpansions
		} else {
			for currentChar := 0; currentChar < len(candidate.data); currentChar++ {
				for possibleLength := 0; possibleLength < maxLengthToAtom; possibleLength++ {
					if currentChar+possibleLength >= len(candidate.data) {
						break
					}
					atomCandidate := candidate.data[currentChar : currentChar+possibleLength+1]
					replacementTos, replacementExists := replacementMap[atomCandidate]
					if !replacementExists {
						continue
					}
					for _, atomReplacement := range replacementTos {
						newNodeData := candidate.data[:currentChar] + atomReplacement + candidate.data[currentChar+possibleLength+1:len(candidate.data)]
						if _, exists := cache[newNodeData]; exists {
							continue
						}
						cache[newNodeData] = true
						newNode := &node{data: newNodeData, numExpansions: candidate.numExpansions + 1}
						for index, candidate := range nodesToTraverse {
							if len(candidate.data) > len(newNodeData) {
								nodesToTraverse = append(nodesToTraverse[:index+1], nodesToTraverse[index:]...)
								nodesToTraverse[index] = *newNode
								break
							}
						}
						nodesToTraverse = append(nodesToTraverse, *newNode)
					}

				}
			}
		}
	}
	return 0
}

func main() {
	inputList := []replacement{}
	var inputString string
	for {
		var inputReplacement replacement
		_, err := fmt.Scanf("%s => %s", &inputReplacement.from, &inputReplacement.to)
		if err != nil {
			break

		}
		inputList = append(inputList, inputReplacement)
	}
	fmt.Scanf("%s", &inputString)
	fmt.Println(partOne(inputList, inputString))
	fmt.Println(partTwo(inputList, inputString))

}
