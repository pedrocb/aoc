package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	parent *node
	childs []*node
	weight int
	name   string
}

var weightCache map[string]int = map[string]int{}

func calcWeight(currentNode *node, nodes map[string]*node) int {
	weight := currentNode.weight
	for _, child := range currentNode.childs {
		if _, exists := weightCache[child.name]; !exists {
			weightCache[child.name] = calcWeight(child, nodes)
		}
		weight += weightCache[child.name]
	}
	return weight
}

func calcIntruderDeviation(weights []int) (int, int) {
	if weights[0] != weights[1] && weights[0] != weights[2] {
		return weights[0] - weights[1], 0
	}
	if weights[1] != weights[2] && weights[1] != weights[0] {
		return weights[1] - weights[2], 1
	}
	if weights[2] != weights[0] && weights[2] != weights[1] {
		return weights[2] - weights[0], 2
	}
	for i := 3; i < len(weights); i++ {
		if weights[i] != weights[i-1] {
			return weights[i-1] - weights[i], i
		}
	}
	return -1, -1
}

func CalcMissingWeight(currentNode *node, nodes map[string]*node) int {
	queue := []*node{currentNode}
	intruderDeviation := 0
	for len(queue) > 0 {
		childWeights := []int{}
		for _, child := range queue[0].childs {
			childWeights = append(childWeights, calcWeight(child, nodes))
		}
		childIntruderDeviation, index := calcIntruderDeviation(childWeights)
		if index != -1 {
			// If intruder found, add childs to queue to calculate their weight
			intruderDeviation = childIntruderDeviation
			queue = append(queue, queue[0].childs[index])
			queue = queue[1:]
		} else {
			// If no child has different weight, we found the uneven node
			return queue[0].weight - intruderDeviation
		}
	}
	return 0
}

func GetParentNode(nodes map[string]*node) *node {
	for _, node := range nodes {
		if node.parent == nil {
			return node
		}
	}
	return nil
}

func main() {
	nodes := map[string]*node{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " -> ")

		nodeInfo := lineSplit[0]
		var nodeName string
		var nodeWeight int
		fmt.Sscanf(nodeInfo, "%s (%d)", &nodeName, &nodeWeight)
		if _, exists := nodes[nodeName]; !exists {
			nodes[nodeName] = &node{name: nodeName}
		}
		nodes[nodeName].weight = nodeWeight

		if len(lineSplit) > 1 {
			childs := strings.Split(lineSplit[1], ", ")
			childNodes := []*node{}
			for _, childName := range childs {
				if _, exists := nodes[childName]; !exists {
					nodes[childName] = &node{name: childName}
				}
				nodes[childName].parent = nodes[nodeName]
				childNodes = append(childNodes, nodes[childName])
			}
			nodes[nodeName].childs = childNodes
		}
	}

	// Part 1
	parentNode := GetParentNode(nodes)
	result1 := parentNode.name
	fmt.Println(result1)

	// Part 2
	result2 := CalcMissingWeight(parentNode, nodes)
	fmt.Println(result2)
}
