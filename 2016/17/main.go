package main

import (
	"crypto/md5"
	"fmt"
)

type pos struct {
	x, y int
}

type node struct {
	passcode   string
	path       string
	currentPos pos
}

// Map characters to directions
var directions map[byte]pos = map[byte]pos{
	'U': pos{x: 0, y: -1},
	'D': pos{x: 0, y: 1},
	'L': pos{x: -1, y: 0},
	'R': pos{x: 1, y: 0},
}

func getNewNodes(currentNode node) []node {
	// List of possible new nodes from the current one
	newNodes := []node{}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(currentNode.passcode+currentNode.path)))
	// For each possible direction
	for index, direction := range []byte{'U', 'D', 'L', 'R'} {
		// If the door is open
		if int(hash[index]) >= int('b') && int(hash[index]) <= int('f') {
			newX := currentNode.currentPos.x + directions[direction].x
			newY := currentNode.currentPos.y + directions[direction].y
			// If in bounds
			if newX >= 0 && newY >= 0 && newX < 4 && newY < 4 {
				// This move is possible, so add to newNodes list
				newNodes = append(newNodes, node{
					passcode:   currentNode.passcode,
					path:       currentNode.path + string(direction),
					currentPos: pos{x: newX, y: newY},
				})
			}
		}
	}
	return newNodes

}

func calcShortestPath(passcode string) string {
	// Breadth First Search
	nodes := []node{node{passcode: passcode, path: "", currentPos: pos{x: 0, y: 0}}}
	for len(nodes) > 0 {
		currentNode := nodes[0]
		nodes = nodes[1:]

		newNodes := getNewNodes(currentNode)
		for _, candidateNode := range newNodes {
			if candidateNode.currentPos.x == 3 && candidateNode.currentPos.y == 3 {
				// Found target position
				return candidateNode.path
			}
			nodes = append(nodes, candidateNode)
		}
	}
	return "No path found"
}

func calcLongestPathLength(passcode string) int {
	nodes := []node{node{passcode: passcode, path: "", currentPos: pos{x: 0, y: 0}}}
	maxPath := 0
	for len(nodes) > 0 {
		currentNode := nodes[0]
		nodes = nodes[1:]
		newNodes := getNewNodes(currentNode)
		for _, candidateNode := range newNodes {
			if candidateNode.currentPos.x == 3 && candidateNode.currentPos.y == 3 {
				if len(candidateNode.path) > maxPath {
					// Found new longest path
					maxPath = len(candidateNode.path)
				}
			} else {
				nodes = append(nodes, candidateNode)
			}
		}
	}
	return maxPath
}

func main() {
	var passcode string
	fmt.Scanf("%s", &passcode)

	// Part 1
	result1 := calcShortestPath(passcode)
	fmt.Println(result1)

	// Part 2
	result2 := calcLongestPathLength(passcode)
	fmt.Println(result2)
}
