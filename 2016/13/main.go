package main

import (
	"fmt"
)

type node struct {
	numSteps int
	x, y     int
}

const (
	SPACE = 1
	WALL  = 2
)

func positionStatus(x, y int, favoriteNumber int) int {
	sum := x*x + 3*x + 2*x*y + y + y*y
	sum += favoriteNumber
	positiveBits := 0
	for sum > 0 {
		positiveBits += sum & 1
		sum >>= 1
	}
	return positiveBits%2 + 1
}

func generateNewNodes(currentNode node, gridStatus map[string]bool, favoriteNumber int) []node {
	result := []node{}
	for _, direction := range [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}} {
		newX := currentNode.x + direction[0]
		newY := currentNode.y + direction[1]
		if newX < 0 || newY < 0 {
			// Out of bounds
			continue
		}
		newPos := repr(newX, newY)
		_, exists := gridStatus[newPos]
		if !exists {
			newPosStatus := positionStatus(newX, newY, favoriteNumber)
			gridStatus[newPos] = true
			if newPosStatus == SPACE {
				// Only a new node to traverse if it was not traversed yet and it is a free space
				result = append(result, node{x: newX, y: newY, numSteps: currentNode.numSteps + 1})
			}
		}
	}
	return result
}

func repr(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func CalcMinStepsTo(x, y int, favoriteNumber int, limit int) int {
	// BFS Search Queue
	queue := []node{node{numSteps: 0, x: 1, y: 1}}

	// Grid positions not needed to traverse (either already traversed or wall)
	gridStatus := map[string]bool{repr(1, 1): true}
	positionsCount := 0
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if limit > 0 && currentNode.numSteps > limit {
			// If limit is set, return the number of positions traversed (for part 2)
			return positionsCount
		}
		positionsCount++

		newNodes := generateNewNodes(currentNode, gridStatus, favoriteNumber)
		for _, node := range newNodes {
			if node.x == x && node.y == y {
				return node.numSteps
			} else {
				queue = append(queue, node)
			}
		}
	}

	return -1
}


func main() {
	var inputNumber int
	fmt.Scanf("%d", &inputNumber)

	// Part 1
	result1 := CalcMinStepsTo(31, 39, inputNumber, -1)
	fmt.Println(result1)

	// Part 2
	result2 := CalcMinStepsTo(-1, -1, inputNumber, 50)
	fmt.Println(result2)
}
