package main

import (
	"fmt"
)

var N = 100
var STEPS = 100

func partOne(inputMap []bool) int {
	currentMap := &inputMap
	lightsOn := 0
	for i := 0; i < STEPS; i++ {
		newMap := make([]bool, N*N)
		lightsOn = 0
		for currentPos := 0; currentPos < N*N; currentPos++ {
			neighborsOn := 0
			currentX := currentPos % N
			currentY := currentPos / N
			for deltaX := -1; deltaX <= 1; deltaX++ {
				for deltaY := -1; deltaY <= 1; deltaY++ {
					if deltaX == 0 && deltaY == 0 {
						continue
					}
					if currentX+deltaX >= 0 && currentX+deltaX < N && currentY+deltaY >= 0 && currentY+deltaY < N {
						if (*currentMap)[(currentX+deltaX)+((currentY+deltaY)*N)] {
							neighborsOn += 1
						}
					}
				}
			}
			if (*currentMap)[currentX+currentY*N] && (neighborsOn != 2 && neighborsOn != 3) {
				newMap[currentX+currentY*N] = false
			} else if !(*currentMap)[currentX+currentY*N] && (neighborsOn == 3) {
				newMap[currentX+currentY*N] = true
			} else {
				newMap[currentX+currentY*N] = (*currentMap)[currentX+currentY*N]
			}
			if newMap[currentX+currentY*N] {
				lightsOn += 1
			}
		}
		currentMap = &newMap

	}
	return lightsOn
}

func partTwo(inputMap []bool) int {
	currentMap := &inputMap
	lightsOn := 0

	(*currentMap)[0+N*0] = true
	(*currentMap)[0+N*(N-1)] = true
	(*currentMap)[(N-1)+N*0] = true
	(*currentMap)[(N-1)+N*(N-1)] = true
	for i := 0; i < STEPS; i++ {
		newMap := make([]bool, N*N)
		lightsOn = 0
		for currentPos := 0; currentPos < N*N; currentPos++ {
			neighborsOn := 0
			currentX := currentPos % N
			currentY := currentPos / N
			if (currentX == 0 && currentY == 0) || (currentX == 0 && currentY == N-1) || (currentX == N-1 && currentY == 0) || (currentX == N-1 && currentY == N-1) {
				newMap[currentX+currentY*N] = true
				lightsOn += 1
				continue
			}
			for deltaX := -1; deltaX <= 1; deltaX++ {
				for deltaY := -1; deltaY <= 1; deltaY++ {
					if deltaX == 0 && deltaY == 0 {
						continue
					}
					if currentX+deltaX >= 0 && currentX+deltaX < N && currentY+deltaY >= 0 && currentY+deltaY < N {
						if (*currentMap)[(currentX+deltaX)+((currentY+deltaY)*N)] {
							neighborsOn += 1
						}
					}
				}
			}
			if (*currentMap)[currentX+currentY*N] && (neighborsOn < 2 || neighborsOn > 3) {
				newMap[currentX+currentY*N] = false
			} else if !(*currentMap)[currentX+currentY*N] && (neighborsOn == 3) {
				newMap[currentX+currentY*N] = true
			} else {
				newMap[currentX+currentY*N] = (*currentMap)[currentX+currentY*N]
			}
			if newMap[currentX+currentY*N] {
				lightsOn += 1
			}
		}
		currentMap = &newMap
	}
	return lightsOn
}

func main() {
	inputMap := make([]bool, N*N)
	for i := 0; i < N*N; i++ {
		var input byte
		_, err := fmt.Scanf("%c", &input)
		if input == '#' {
			inputMap[i] = true
		} else if input == '\n' {
			i--
		}
		if err != nil {
			break

		}
	}
	fmt.Println(partOne(inputMap))
	fmt.Println(partTwo(inputMap))

}
