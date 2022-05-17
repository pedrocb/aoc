package main

import (
	"fmt"
	"math"
)

func partOne(inputList []int) int {
	numCombinations := math.Pow(2, float64(len(inputList)))
	result := 0
	for candidate := 0; candidate < int(numCombinations); candidate++ {
		counter := 0
		for j := 0; j < len(inputList); j++ {
			if candidate>>j&1 == 1 {
				counter += inputList[j]
				if counter > 150 {
					break
				}
			}
		}
		if counter == 150 {
			result += 1
		}
	}
	return result
}

func partTwo(inputList []int) int {
	numCombinations := math.Pow(2, float64(len(inputList)))
	result := 0
	minimumContainers := len(inputList)
	for candidate := 0; candidate < int(numCombinations); candidate++ {
		counter := 0
		bitCounter := 0
		for j := 0; j < len(inputList); j++ {
			if candidate>>j&1 == 1 {
				counter += inputList[j]
				bitCounter += 1
				if counter > 150 || bitCounter > minimumContainers {
					break
				}
			}
		}
		if counter == 150 {
			if bitCounter < minimumContainers {
				minimumContainers = bitCounter
				result = 1
			} else if bitCounter == minimumContainers {
				result += 1
			}
		}
	}
	return result
}

func main() {
	var inputList []int
	for {
		var input int
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			break

		}
		inputList = append(inputList, input)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))

}
