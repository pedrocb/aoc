package main

import (
	"fmt"
	"github.com/pedrocb/aoc/common/assembunny"
)

func FindLowestIntegerLoop(inputInstructions []assembunny.Instruction, confidenceSize int, loopInterval []int) int {
	for candidate := 0; true; candidate++ {
		if assembunny.FindOutputLoop(inputInstructions, map[string]int{
			"a": candidate,
			"b": 0,
			"c": 0,
			"d": 0,
		}, confidenceSize, loopInterval) {
			return candidate
		}
	}
	return -1
}

func main() {
	inputInstructions := assembunny.ScanInstructions()

	// Part One
	result1 := FindLowestIntegerLoop(inputInstructions, 10, []int{0, 1})
	fmt.Println(result1)

}
