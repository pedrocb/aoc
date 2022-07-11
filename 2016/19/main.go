package main

import (
	"fmt"
)

func calcWinner(numElves int) int {
	// Josephus Problem
	if numElves == 1 {
		// If only one elf, the first one wins
		return 1
	}
	// Otherwise, it is the one to the left of the winner with one less elf (calcWinner(numElves-1) + 1)%numElves
	return (calcWinner(numElves-1)+1)%numElves + 1
}

func calcWinnerAcrossCircle(numElves int) int {
	// For efficiency keep a queue for each side of queue
	// This way we only need to delete/insert in the first/last position
	leftQueue := []int{}
	rightQueue := []int{}
	for elf := 0; elf < numElves; elf++ {
		//  Left side to left queue and vice versa.
		if elf < (numElves+1)/2 {
			leftQueue = append(leftQueue, elf)
		} else {
			rightQueue = append(rightQueue, elf)
		}
	}
	circleLength := numElves
	// When we reach a circle with size 2, the winner is the next elf to steal
	for circleLength > 2 {
		// Current elf is always the first on left queue
		currentElf := leftQueue[0]

		if circleLength%2 != 0 {
			// If the circle is odd, neighbor is the last in left queue
			leftQueue = leftQueue[:len(leftQueue)-1]
		} else {
			// Otherwise is the first on the right queue
			rightQueue = rightQueue[1:]
		}
		// Rotate queues, so the first of right queue is moved to last of left queue
		// and first of left queue is moved to last of right queue (1, 2),(3, 4) -> (2, 3),(4,1)
		leftQueue = leftQueue[1:]
		leftQueue = append(leftQueue, rightQueue[0])
		rightQueue = append(rightQueue[1:], currentElf)
		// Discount removed elf
		circleLength--
	}
	return leftQueue[0] + 1
}

func main() {
	var numElves int
	fmt.Scanf("%d", &numElves)

	// Part 1
	result1 := calcWinner(numElves)
	fmt.Println(result1)

	// Part 2
	result2 := calcWinnerAcrossCircle(numElves)
	fmt.Println(result2)
}
