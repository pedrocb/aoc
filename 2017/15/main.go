package main

import (
	"fmt"
	"math"
)

var genAFactor int = 16807
var genBFactor int = 48271

func CountMatches(genASeed, genBSeed, nValues, multipleA, multipleB int) int {
	counter := 0
	currentA := genASeed
	currentB := genBSeed
	mask := int(math.Pow(2, 16)) - 1
	for i := 0; i < nValues; i++ {
		for {
			currentA = (currentA * genAFactor) % 2147483647
			if currentA%multipleA == 0 {
				break
			}
		}
		for {
			currentB = (currentB * genBFactor) % 2147483647
			if currentB%multipleB == 0 {
				break
			}
		}
		if currentA&mask == currentB&mask {
			counter++
		}
	}
	return counter
}

func main() {
	var genASeed, genBSeed int
	fmt.Scanf("Generator A starts with %d", &genASeed)
	fmt.Scanf("Generator B starts with %d", &genBSeed)

	// Part 1
	result1 := CountMatches(genASeed, genBSeed, 40000000, 1, 1)
	fmt.Println(result1)

	// Part 2
	result2 := CountMatches(genASeed, genBSeed, 5000000, 4, 8)
	fmt.Println(result2)
}
