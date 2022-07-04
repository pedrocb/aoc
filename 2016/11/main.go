package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	GEN  = "generator"
	CHIP = "microchip"
)

type state struct {
	components []component
	elevator   int
}

type component struct {
	gen, chip int
}

type node struct {
	numSteps  int
	state     state
	heuristic int
	index     int
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].numSteps+pq[i].heuristic < pq[j].numSteps+pq[j].heuristic
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = i
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func normalizeName(rawComponentName string, componentType string) string {
	if componentType == CHIP {
		return rawComponentName[:len(rawComponentName)-len("-compatible")]
	}
	return rawComponentName
}

func addComponent(elements map[string]component, componentName string, componentType string, floor int) {
	componentObject, exists := elements[componentName]
	if !exists {
		componentObject = component{}
	}

	if componentType == GEN {
		componentObject.gen = floor
	} else {
		componentObject.chip = floor
	}
	elements[componentName] = componentObject
}

func validFloor(state state, floor int) bool {
	foundChipAlone := false
	foundGenerators := false
	for _, component := range state.components {
		if component.chip == floor && component.gen == floor {
			foundGenerators = true
		} else if component.gen == floor && component.chip != floor {
			foundGenerators = true
		} else if component.chip == floor && component.gen != floor {
			foundChipAlone = true
		}
		if foundChipAlone && foundGenerators {
			// Floor is invalid if there is a chip with other generators other than its own
			return false
		}
	}
	return true
}

func calcHeuristic(state state) int {
	// Heuristic is how many steps it would take if there weren't any restrictions and each elevator would move 2 items up
	counter := 0
	for _, component := range state.components {
		counter += 3 - component.chip
		counter += 3 - component.gen
	}
	return counter / 2
}

func move(currentState state, direction int, componentIndex int) (state, bool) {
	resultState := state{
		components: make([]component, len(currentState.components)),
		elevator:   currentState.elevator + direction,
	}
	copy(resultState.components, currentState.components)
	if componentIndex%2 == 0 {
		resultState.components[componentIndex/2].gen += direction
	} else {
		resultState.components[componentIndex/2].chip += direction
	}

	// Return resultState, and if it is valid or not
	// State is valid if both affected floors are valid
	return resultState, validFloor(resultState, currentState.elevator) && validFloor(resultState, resultState.elevator)
}

func generateNewNodes(currentNode node) []node {
	currentState := currentNode.state
	numComponents := len(currentNode.state.components)
	newNodes := []node{}

	for firstComponentIndex := 0; firstComponentIndex < numComponents*2; firstComponentIndex++ {
		var componentFloor int
		if firstComponentIndex%2 == 0 {
			componentFloor = currentState.components[firstComponentIndex/2].gen
		} else {
			componentFloor = currentState.components[firstComponentIndex/2].chip
		}
		if componentFloor != currentState.elevator {
			// If component not on elevator floor, it can't be moved
			continue
		}
		for secondComponentIndex := firstComponentIndex; secondComponentIndex < numComponents*2; secondComponentIndex++ {
			// Both up and down directions
			for _, direction := range []int{-1, 1} {

				// If in bounds
				if currentState.elevator+direction >= 0 && currentState.elevator+direction <= 3 {

					// Calculate moving only first component
					newState, validState := move(currentState, direction, firstComponentIndex)
					if secondComponentIndex == firstComponentIndex {
						// If both components have the same index, we only move the one
						if validState {
							newNodes = append(newNodes, node{
								numSteps:  currentNode.numSteps + 1,
								state:     newState,
								heuristic: 0,
							})
						} else {
							continue
						}
					}

					// If the second component can't be moved continue
					if secondComponentIndex%2 == 0 {
						componentFloor = currentState.components[secondComponentIndex/2].gen
					} else {
						componentFloor = currentState.components[secondComponentIndex/2].chip
					}
					if componentFloor != currentState.elevator {
						continue
					}

					// Both are of the same type or both are of the same element
					if firstComponentIndex%2 != secondComponentIndex%2 && (firstComponentIndex/2 != secondComponentIndex/2) {
						continue
					}

					// Otherwise move both components
					newState.elevator = currentState.elevator
					newState, validState = move(newState, direction, secondComponentIndex)
					if validState {
						newNodes = append(newNodes, node{
							numSteps:  currentNode.numSteps + 1,
							state:     newState,
							heuristic: 0,
						})
					}
				}
			}

		}
	}

	return newNodes

}

func isFinalState(inputState state) bool {
	// Is final state if all components are on 4th floor
	for _, component := range inputState.components {
		if component.gen != 3 || component.chip != 3 {
			return false
		}
	}
	return true
}

func repr(inputState state) string {
	// State representation in string so it can be cached

	// First sort, then convert to string
	sort.Slice(inputState.components, func(i, j int) bool {
		if inputState.components[i].gen == inputState.components[j].gen {
			return inputState.components[i].chip < inputState.components[j].chip
		}
		return inputState.components[i].gen < inputState.components[j].gen
	})
	return fmt.Sprintf("%v-%d", inputState.components, inputState.elevator)
}

func CalcMinSteps(initialState state) int {
	// A* algorithm
	cache := map[string]bool{}

	// Queue with states
	pq := make(PriorityQueue, 1)
	pq[0] = &node{numSteps: 0, state: initialState, heuristic: 0}
	heap.Init(&pq)

	for pq.Len() > 0 {
		// Get node with most priority (< numStepsTake + heuristic optimal steps missing)
		currentNode := heap.Pop(&pq).(*node)

		// Generate all new valid possible states
		newNodes := generateNewNodes(*currentNode)

		for index, candidateNode := range newNodes {
			if isFinalState(candidateNode.state) {
				// FInal condition found
				return candidateNode.numSteps
			}

			// If state is not cached, add to queue and cache it
			stateRepr := repr(candidateNode.state)
			if _, exists := cache[stateRepr]; !exists {
				candidateNode.heuristic = calcHeuristic(candidateNode.state)

				heap.Push(&pq, &newNodes[index])
				cache[stateRepr] = true
			}
		}
	}
	return -1
}

func main() {
	bio := bufio.NewReader(os.Stdin)
	currentFloor := 0

	// Elements arrangement representation
	// Key is the element name and value is [gen floor, chip floor]
	elementsMap := map[string]component{}
	for {
		// Read line
		inputString, err := bio.ReadString('\n')
		inputString = strings.TrimSuffix(inputString, "\n")
		if err != nil {
			break
		}

		splitLine := strings.Split(inputString, "and a ")
		if len(splitLine) > 1 {
			// If line has "and a", the next words are the last element
			lastElement := strings.TrimSuffix(splitLine[1], ".")

			var lastElementName, lastElementType string
			fmt.Sscanf(lastElement, "%s %s", &lastElementName, &lastElementType)
			lastElementName = normalizeName(lastElementName, lastElementType)
			addComponent(elementsMap, lastElementName, lastElementType, currentFloor)
		}

		// All the other elements are separated by comma
		otherElements := strings.Split(splitLine[0], ",")
		for index, element := range otherElements {
			var floor, elementName, elementType string
			if index == 0 {
				fmt.Sscanf(element, "The %s floor contains a %s %s", &floor, &elementName, &elementType)
			} else {
				fmt.Sscanf(element, " a %s %s", &elementName, &elementType)
			}
			if elementName != "" {
				elementName = normalizeName(elementName, elementType)
				addComponent(elementsMap, elementName, elementType, currentFloor)
			}
		}
		currentFloor++
	}

	// Convert map to list of value since the name of the elements is not relevant, only the positions
	components := []component{}
	for _, value := range elementsMap {
		components = append(components, value)
	}
	result1 := CalcMinSteps(state{components: components, elevator: 0})
	fmt.Println(result1)

	components = append(components, component{gen: 0, chip: 0})
	components = append(components, component{gen: 0, chip: 0})
	result2 := CalcMinSteps(state{components: components, elevator: 0})
	fmt.Println(result2)
}
