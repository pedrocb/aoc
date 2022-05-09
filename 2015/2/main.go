package main

import (
	"fmt"
	"sort"
)

type box struct {
	length int
	width  int
	height int
}

func (b *box) sortedDimensions() []int {
	dimensions := []int{b.length, b.width, b.height}
	sort.Ints(dimensions)
	return dimensions
}

func partOne(inputBoxes []box) int {
	totalArea := 0
	for _, inputBox := range inputBoxes {
		sortedDimensions := inputBox.sortedDimensions()
		minimumSide := sortedDimensions[0] * sortedDimensions[1]
		totalArea += 3*minimumSide +
			2*sortedDimensions[1]*sortedDimensions[2] +
			2*sortedDimensions[0]*sortedDimensions[2]
	}
	return totalArea
}

func partTwo(inputBoxes []box) int {
	totalRibbon := 0
	for _, inputBox := range inputBoxes {
		sortedDimensions := inputBox.sortedDimensions()
		totalRibbon += sortedDimensions[0]*2 + sortedDimensions[1]*2
		totalRibbon += sortedDimensions[0] * sortedDimensions[1] * sortedDimensions[2]
	}
	return totalRibbon
}

func main() {
	var inputBox box
	inputBoxes := []box{}
	for {
		_, err := fmt.Scanf("%dx%dx%d", &inputBox.length, &inputBox.width, &inputBox.height)
		if err != nil {
			break
		}
		inputBoxes = append(inputBoxes, inputBox)
	}
	fmt.Println(partOne(inputBoxes))
	fmt.Println(partTwo(inputBoxes))
}
