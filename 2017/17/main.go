package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

func Spinlock(steps, n int) *node {
	currentNode := &node{value: 0}
	currentNode.next = currentNode
	for value := 1; value <= n; value++ {
		for step := 0; step < steps; step++ {
			currentNode = currentNode.next
		}
		nextValue := currentNode.next
		currentNode.next = &node{value: value, next: nextValue}
		currentNode = currentNode.next
	}
	return currentNode
}

func SpinlockAfter0(steps, n int) int {
	result := -1
	curPos := 0
	for i := 1; i <= n; i++ {
		curPos = (curPos + steps) % i
		if curPos == 0 {
			result = i
		}
		curPos++
	}
	return result
}

func main() {
	var steps int
	fmt.Scanf("%d", &steps)

	// Part 1
	result1 := Spinlock(steps, 2017)
	fmt.Println(result1.next.value)

	// Part 2
	result2 := SpinlockAfter0(steps, 50000000)
	fmt.Println(result2)
}
