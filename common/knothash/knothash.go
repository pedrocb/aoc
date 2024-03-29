package knothash

import (
	"fmt"
	"strings"
)

var standardLengthSuffixValues []int = []int{17, 31, 73, 47, 23}
var N int = 256

func toAscii(input string) []int {
	result := []int{}
	for _, char := range input {
		result = append(result, int(char))
	}
	return result
}

func reduceHash(hash []int) []int {
	result := make([]int, N/16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			result[i] ^= hash[i*16+j]
		}
	}
	return result
}
func toHex(input []int) string {
	var result strings.Builder
	for _, number := range input {
		result.WriteString(fmt.Sprintf("%02x", number))
	}
	return result.String()
}

func sequence(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i
	}
	return result
}

func KnotHashInternals(lengths []int, rounds int) []int {
	result := sequence(N)
	currentPosition := 0
	skipSize := 0
	for round := 0; round < rounds; round++ {
		for _, length := range lengths {
			start := currentPosition
			end := (currentPosition + length - 1) % N
			for i := 0; i < length/2; i++ {
				temp := result[start]
				result[start] = result[end]
				result[end] = temp
				start++
				end--
				if start == N {
					start = 0
				}
				if end < 0 {
					end = N - 1
				}
			}
			currentPosition = (currentPosition + length + skipSize) % N
			skipSize++
		}
	}

	return result
}

func KnotHash(inputString string) string {
	asciiLengths := toAscii(inputString)
	asciiLengths = append(asciiLengths, standardLengthSuffixValues...)
	return toHex(reduceHash(KnotHashInternals(asciiLengths, 64)))
}
