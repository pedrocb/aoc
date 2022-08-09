package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id         int
	neighbours []*node
}

func FindConnectedNodes(initialNode *node) []*node {
	nodesVisited := map[int]bool{initialNode.id: true}
	queue := []*node{initialNode}
	result := []*node{initialNode}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, neighbour := range currentNode.neighbours {
			if _, exists := nodesVisited[neighbour.id]; !exists {
				queue = append(queue, neighbour)
				result = append(result, neighbour)
				nodesVisited[neighbour.id] = true
			}
		}
	}
	return result
}

func FindGroups(nodes []*node) [][]*node {
	nodesVisited := map[int]bool{}
	result := [][]*node{}
	for _, node := range nodes {
		if _, exists := nodesVisited[node.id]; !exists {
			connectedNodes := FindConnectedNodes(node)
			for _, connectedNode := range connectedNodes {
				nodesVisited[connectedNode.id] = true
			}
			result = append(result, connectedNodes)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nodes := map[int]*node{}
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), ", ")
		var nodeId int
		var firstNeighbour string
		fmt.Sscanf(lineSplit[0], "%d <-> %s", &nodeId, &firstNeighbour)
		inputNode, exists := nodes[nodeId]
		if !exists {
			inputNode = &node{id: nodeId, neighbours: []*node{}}
		}
		neighbours := append([]string{firstNeighbour}, lineSplit[1:]...)
		for _, neighbour := range neighbours {
			neighbourId, _ := strconv.Atoi(neighbour)
			neighbourNode, exists := nodes[neighbourId]
			if !exists {
				neighbourNode = &node{id: neighbourId, neighbours: []*node{}}
			}
			inputNode.neighbours = append(inputNode.neighbours, neighbourNode)
			nodes[neighbourId] = neighbourNode
		}
		nodes[nodeId] = inputNode
	}
	result1 := len(FindConnectedNodes(nodes[0]))
	fmt.Println(result1)

	nodesList := []*node{}
	for _, n := range nodes {
		nodesList = append(nodesList, n)
	}
	result2 := len(FindGroups(nodesList))
	fmt.Println(result2)
}
