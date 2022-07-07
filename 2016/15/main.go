package main

import (
	"fmt"
)

type disc struct {
	numPositions    int
	currentPosition int
}

func positionsAt(discs []disc, time int) []int {
	// Get all disc positions at time time
	positions := []int{}
	for index, currentDisc := range discs {
		positions = append(positions, (currentDisc.currentPosition+(index+1)+time)%currentDisc.numPositions)
	}

	return positions
}

func firstTimeToPressButton(discs []disc) int {
	candidateTime := 0
	for {
		positions := positionsAt(discs, candidateTime)
		getsCapsule := true
		for _, pos := range positions {
			// Only when all discs are at 0 that the capsule falls through
			if pos != 0 {
				getsCapsule = false
				break
			}
		}
		if getsCapsule {
			return candidateTime
		}
		candidateTime++
	}
}

func main() {
	inputDiscs := []disc{}
	for {
		var inputDisc disc
		var discNumber string
		_, err := fmt.Scanf("Disc %s has %d positions; at time=0, it is at position %d.\n", &discNumber, &inputDisc.numPositions, &inputDisc.currentPosition)
		if err != nil {
			break
		}
		inputDiscs = append(inputDiscs, inputDisc)
	}

	// Part 1
	fmt.Println(firstTimeToPressButton(inputDiscs))

	// Part 2
	inputDiscs = append(inputDiscs, disc{numPositions: 11, currentPosition: 0})
	fmt.Println(firstTimeToPressButton(inputDiscs))
}
