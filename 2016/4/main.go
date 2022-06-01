package main

import (
	"fmt"
	"sort"
	"strings"
)

type room struct {
	name     string
	id       int
	checksum string
}

type keyValue struct {
	key   rune
	value int
}

func calcChecksum(inputRoom room) string {
	// Calculate distribution of chars
	charDistribution := map[rune]int{}
	for _, char := range inputRoom.name {
		if char != '-' {
			charDistribution[char] += 1
		}
	}

	// Create a slice with distributions, and sort it based on the occurrence, and alphabeticaly
	charDistributionSlice := []keyValue{}
	for k, v := range charDistribution {
		charDistributionSlice = append(charDistributionSlice, keyValue{key: k, value: v})
	}
	sort.Slice(charDistributionSlice, func(l, r int) bool {
		if charDistributionSlice[l].value == charDistributionSlice[r].value {
			return int(charDistributionSlice[l].key) < int(charDistributionSlice[r].key)
		}
		return charDistributionSlice[l].value > charDistributionSlice[r].value
	})

	// Build string with first most common chars
	var checksum strings.Builder
	for _, kv := range charDistributionSlice[:5] {
		checksum.WriteRune(kv.key)
	}

	return checksum.String()
}

func decryptRoomName(inputRoom room) string {
	var decryptedName strings.Builder
	alphabetSize := int('z') - int('a') + 1
	for _, char := range inputRoom.name {
		decryptedChar := ' '
		if char != '-' {
			decryptedChar = rune(((int(char) - int('a') + inputRoom.id) % alphabetSize) + int('a'))

		}
		decryptedName.WriteRune(decryptedChar)
	}

	return decryptedName.String()
}

func CalcIdSumRealRooms(inputRooms []room) int {
	sum := 0

	for _, currentRoom := range inputRooms {
		checksum := calcChecksum(currentRoom)
		if checksum == currentRoom.checksum {
			sum += currentRoom.id
		}
	}

	return sum
}

func FindNorthpoleStorageRoomId(inputRooms []room) int {
	// Find room with word northpole in it
	for _, inputRoom := range inputRooms {
		decryptedName := decryptRoomName(inputRoom)
		for _, word := range strings.Split(decryptedName, " ") {
			if word == "northpole" {
				return inputRoom.id
			}
		}
	}
	return -1
}

func main() {
	var inputRooms []room
	for {
		var inputRoom room
		var inputString string

		// Input format: name_of_room-id[checksum]
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
		// Split string by -, and scan the last part id[checksum]
		splitString := strings.Split(inputString, "-")
		lastSection := splitString[len(splitString)-1]
		fmt.Sscanf(lastSection[:len(lastSection)-1], "%d[%s", &inputRoom.id, &inputRoom.checksum)

		// The rest is the room name
		inputRoom.name = strings.Join(splitString[:len(splitString)-1], "-")
		inputRooms = append(inputRooms, inputRoom)
	}

	// Puzzle 1
	result1 := CalcIdSumRealRooms(inputRooms)
	fmt.Println(result1)

	// Puzzle 2
	result2 := FindNorthpoleStorageRoomId(inputRooms)
	fmt.Println(result2)
}
