package main

import (
	"fmt"
	"math"
	"strconv"
)

var width int = 179
var height int = 39
var numLocations int = 7

type node struct {
	currentX, currentY int
	visitedLocations   int
	numSteps           int
}

func (n *node) setLocation(location int) {
	n.visitedLocations |= (1 << (location - 1))
}

func (n node) repr() string {
	return fmt.Sprintf("%d-%d-%d", n.currentX, n.currentY, n.visitedLocations)
}

func (n node) foundLocations(grid [][]byte) bool {
	return n.visitedLocations == int(math.Pow(2, float64(numLocations))-1)
}

func (n node) foundLocationsAndReturned(grid [][]byte) bool {
	return n.foundLocations(grid) && grid[n.currentX][n.currentY] == '0'
}

func (n node) getNewNodes(grid [][]byte) []node {
	newNodes := []node{}
	for _, direction := range [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}} {
		newX := n.currentX + direction[0]
		newY := n.currentY + direction[1]
		switch grid[newX][newY] {
		case '#':
			continue
		case '.', '0':
			newNodes = append(newNodes, node{currentX: newX, currentY: newY, visitedLocations: n.visitedLocations, numSteps: n.numSteps + 1})
		default:
			location, _ := strconv.Atoi(string(grid[newX][newY]))
			newNode := node{currentX: newX, currentY: newY, visitedLocations: n.visitedLocations, numSteps: n.numSteps + 1}
			newNode.setLocation(location)
			newNodes = append(newNodes, newNode)
		}
	}
	return newNodes
}

func calcShortestPath(grid [][]byte, initialX, initialY int, finalMethod func(n node, grid [][]byte) bool) int {
	// BFS algorithm
	tree := []node{node{currentX: initialX, currentY: initialY, visitedLocations: 0, numSteps: 0}}
	cache := map[string]bool{}
	for len(tree) > 0 {
		currentNode := tree[0]
		tree = tree[1:]

		newNodes := currentNode.getNewNodes(grid)
		for _, node := range newNodes {
			// Logic is the same for both parts, only the final state is different, so the method to calculate
			// the final state is passed as argument
			if finalMethod(node, grid) {
				return node.numSteps
			}
			if _, exists := cache[node.repr()]; !exists {
				tree = append(tree, node)
				cache[node.repr()] = true
			}
		}

	}
	return -1
}

func main() {
	grid := make([][]byte, width)
	var initialX, initialY int
	for y := 0; y < height; y++ {
		for x := 0; x <= width; x++ {
			if x == width {
				fmt.Scanf("\n")
				continue
			}

			if y == 0 {
				grid[x] = make([]byte, height)
			}

			fmt.Scanf("%c", &grid[x][y])
			if grid[x][y] == '0' {
				initialX = x
				initialY = y
			}
		}
	}

	// Part 1
	result1 := calcShortestPath(grid, initialX, initialY, node.foundLocations)
	fmt.Println(result1)

	// Part 2
	result2 := calcShortestPath(grid, initialX, initialY, node.foundLocationsAndReturned)
	fmt.Println(result2)
}
