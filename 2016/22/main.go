package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	x, y       int
	size, used int
}

func (n node) free() int {
	return n.size - n.used
}

type treeNode struct {
	emptyNodeX, emptyNodeY int
	goalDataX, goalDataY   int
	numSteps               int
}

type pair struct {
	first, second node
}

var NCol int = 32
var NRow int = 31

func flattenGrid(grid [][]node) []node {
	result := []node{}
	for x := 0; x < NCol; x++ {
		result = append(result, grid[x]...)
	}
	return result
}

func CalcViablePairs(grid [][]node) []pair {
	// Sort nodes by free size
	nodes := flattenGrid(grid)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].free() > nodes[j].free()
	})
	viablePairs := []pair{}
	for currentNodeIndex, currentNode := range nodes {
		// Ignore empty node
		if currentNode.used <= 0 {
			continue
		}
		for pairNodeIndex, pairNode := range nodes {
			if currentNodeIndex == pairNodeIndex {
				// Same node
				continue
			}
			if currentNode.used <= pairNode.free() {
				viablePairs = append(viablePairs, pair{first: currentNode, second: pairNode})
			} else {
				break
			}
		}

	}
	return viablePairs
}

func isOutOfBounds(x, y int) bool {
	return x < 0 || x >= NCol || y < 0 || y >= NRow
}

func getNewNodes(grid [][]node, currentNode treeNode) []treeNode {
	newNodes := []treeNode{}

	for _, direction := range [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}} {
		newX := currentNode.emptyNodeX + direction[0]
		newY := currentNode.emptyNodeY + direction[1]
		if isOutOfBounds(newX, newY) {
			continue
		}
		if grid[newX][newY].used > grid[currentNode.emptyNodeX][currentNode.emptyNodeY].size {
			continue
		}
		goalDataX := currentNode.goalDataX
		goalDataY := currentNode.goalDataY
		if newX == goalDataX && newY == goalDataY {
			goalDataX = currentNode.emptyNodeX
			goalDataY = currentNode.emptyNodeY
		}

		newNodes = append(newNodes, treeNode{
			goalDataX:  goalDataX,
			goalDataY:  goalDataY,
			emptyNodeX: newX,
			emptyNodeY: newY,
			numSteps:   currentNode.numSteps + 1,
		})
	}
	return newNodes
}

func repr(input treeNode) string {
	return fmt.Sprintf("%d-%d-%d-%d", input.emptyNodeX, input.emptyNodeY, input.goalDataX, input.goalDataY)
}

func CalcMinimumNumberSteps(grid [][]node) int {
	// BFS algorithm
	// There is only one empty node that can receive data
	tree := []treeNode{treeNode{emptyNodeX: 28, emptyNodeY: 20, numSteps: 0, goalDataX: NCol - 1, goalDataY: 0}}
	cache := map[string]bool{}
	for {
		currentNode := tree[0]
		tree = tree[1:]

		newNodes := getNewNodes(grid, currentNode)
		for _, newNode := range newNodes {
			if newNode.goalDataX == 0 && newNode.goalDataY == 0 {
				// Found final state
				return newNode.numSteps
			}
			if _, exists := cache[repr(newNode)]; !exists {
				// If cached, already processed
				tree = append(tree, newNode)
				cache[repr(newNode)] = true
			}
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Ignore first two lines
	scanner.Scan()
	scanner.Scan()

	grid := make([][]node, NCol)
	for scanner.Scan() {
		var inputNode node
		var nodeX, nodeY int
		_, err := fmt.Sscanf(scanner.Text(), "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%", &nodeX, &nodeY, &inputNode.size, &inputNode.used, new(int), new(int))
		if err != nil {
			fmt.Println(err)
			return
		}
		if nodeY == 0 {
			grid[nodeX] = make([]node, NRow)
		}
		grid[nodeX][nodeY] = inputNode
	}

	// Part 1
	viablePairs := CalcViablePairs(grid)
	fmt.Println(len(viablePairs))

	// Part 2
	result2 := CalcMinimumNumberSteps(grid)
	fmt.Println(result2)
}
