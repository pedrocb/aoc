package main

import (
	"fmt"
	"strings"
)

func modifiedDragonCurve(a []bool) []bool {
	b := []bool{false}
	for index := len(a) - 1; index >= 0; index-- {
		b = append(b, !a[index])
	}
	return append(a, b...)
}

func calcChecksum(data []bool) []bool {
	dataSize := len(data)

	// Calc size of checksum
	numChunks := dataSize
	for numChunks%2 == 0 {
		numChunks /= 2
	}

	chunkSize := dataSize / numChunks
	result := []bool{}
	for currentChunk := 0; currentChunk < numChunks; currentChunk++ {
		// Each chunk corresponds to each bit of the checksum
		// If chunk has pair number of 1s the checksuum bit is 1 otherwise is 0
		chunkBit := true
		for bit := 0; bit < chunkSize; bit++ {
			if data[currentChunk*chunkSize+bit] {
				chunkBit = !chunkBit
			}
		}
		result = append(result, chunkBit)
	}
	return result
}

func generateData(size int, initialState []bool) ([]bool, []bool) {
	data := initialState
	for len(data) < size {
		data = modifiedDragonCurve(data)
	}
	data = data[:size]
	checksum := calcChecksum(data)
	return data, checksum
}

func repr(input []bool) string {
	var result strings.Builder
	for _, bit := range input {
		if bit {
			result.WriteByte('1')
		} else {
			result.WriteByte('0')
		}
	}
	return result.String()
}

func main() {
	// Represent Data as slice of bools
	initialState := []bool{}
	for {
		var inputChar byte
		_, err := fmt.Scanf("%c", &inputChar)
		if err != nil {
			break
		}
		if inputChar == '1' {
			initialState = append(initialState, true)
		} else if inputChar == '0' {
			initialState = append(initialState, false)
		}
	}

	// Part 1
	_, result1 := generateData(272, initialState)
	fmt.Println(repr(result1))

	// Part 2
	_, result2 := generateData(35651584, initialState)
	fmt.Println(repr(result2))
}
