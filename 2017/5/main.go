package main

import (
	"fmt"
)

func CalculateSteps(inputJumps []int, tuneJump func(int) int) []int {
	jumps := make([]int, len(inputJumps))
	steps := []int{}
	copy(jumps, inputJumps)

	for index := 0; index < len(jumps); {
		steps = append(steps, jumps[index])
		jumps[index] = tuneJump(jumps[index])
		index += jumps[index] - 1
	}
	return steps
}

func main() {
	jumps := []int{}
	for {
		var inputJump int
		_, err := fmt.Scanf("%d", &inputJump)
		if err != nil {
			break
		}
		jumps = append(jumps, inputJump)
	}

	// Part 1
	result1 := len(CalculateSteps(jumps, func(i int) int { return i + 1 }))
	fmt.Println(result1)

	// Part 2
	result2 := len(CalculateSteps(jumps, func(i int) int {
		if i > 3 {
			return i - 1
		} else {
			return i + 1
		}
	}))
	fmt.Println(result2)
}
