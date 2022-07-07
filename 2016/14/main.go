package main

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
)

type key struct {
	value     string
	saltIndex int
	triplet   byte
}

func GetSequences(inputString string, minSequenceSize int) []string {
	sequences := []string{}
	currentSequenceChar := inputString[0]
	currentSequenceSize := 1
	for index := 1; index < len(inputString); index++ {
		if inputString[index] == currentSequenceChar {
			currentSequenceSize++
		} else {
			if currentSequenceSize >= minSequenceSize {
				sequences = append(sequences, inputString[index-currentSequenceSize:index])
			}
			currentSequenceChar = inputString[index]
			currentSequenceSize = 1
		}
	}
	if currentSequenceSize >= minSequenceSize {
		sequences = append(sequences, inputString[len(inputString)-currentSequenceSize:len(inputString)])
	}
	return sequences
}

func GetNKeys(n int, salt string, nHashings int) []key {
	saltIndex := 0

	// Keys with triplets but 5ths not found yet
	candidateKeys := []key{}
	// Valid keys
	resultKeys := []key{}
	// Max saltIndex of valid keys
	maxKey := -1

	// We can't stop when n valid keys are found because lower candidate keys can be found afterwards
	// While there are candidate keys with lower saltIndex then the last valid key, keep generating
	for len(resultKeys) < n || candidateKeys[0].saltIndex < maxKey {
		composedSalt := salt + strconv.Itoa(saltIndex)
		// Key stretching
		hash := composedSalt
		for i := 0; i < nHashings; i++ {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
		}

		// Get all sequences > 3
		sequences := GetSequences(hash, 3)

		foundTriplet := false
		for _, sequence := range sequences {
			// If first triplet found, added key to the candidate keys
			if !foundTriplet && len(sequence) >= 3 {
				candidateKeys = append(candidateKeys, key{value: hash, saltIndex: saltIndex, triplet: sequence[0]})
				foundTriplet = true
			}
			// If fifthlet found, check if validates any candidate key
			if len(sequence) >= 5 {
				for index := 0; index < len(candidateKeys) && candidateKeys[index].saltIndex < saltIndex; index++ {
					if candidateKeys[index].triplet == sequence[0] {
						// If it validates, add key to result keys, update maxKey value and remove key from candidate keys
						resultKeys = append(resultKeys, candidateKeys[index])

						if maxKey == -1 || maxKey < candidateKeys[index].saltIndex {
							maxKey = candidateKeys[index].saltIndex
						}
						candidateKeys = append(candidateKeys[:index], candidateKeys[index+1:]...)
						index--
					}
				}
			}
		}

		// Expire all candidate that no longer can be validated since all the 1000 next keys were tested
		numCandidatesToRemove := 0
		for _, candidateKey := range candidateKeys {
			if candidateKey.saltIndex <= saltIndex-1000 {
				numCandidatesToRemove++
			} else {
				break
			}
		}
		candidateKeys = candidateKeys[numCandidatesToRemove:]

		saltIndex++

	}

	// Sort valid keys by salt Index, since they are not found in order
	sort.Slice(resultKeys, func(i, j int) bool { return resultKeys[i].saltIndex < resultKeys[j].saltIndex })
	return resultKeys[:n]
}

func main() {
	var salt string
	fmt.Scanf("%s", &salt)

	// Part 1
	result1 := GetNKeys(64, salt, 1)
	fmt.Println(result1[63].saltIndex)

	// Part 2
	result2 := GetNKeys(64, salt, 2017)
	fmt.Println(result2[63].saltIndex)
}
