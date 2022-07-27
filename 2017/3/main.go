package main

import (
	"fmt"
)

func CalcShortestPath(square int) int {
	if square == 1 {
		return 0
	}
	// Size of current square
	squareSide := 1
	// MinValue in current square
	minSquare := 1
	// MaxValue in current square
	maxSquare := 1

	// Calculate min and max value of each innersquare
	// If square with value found, calculate distance based on distance to corner
	for {
		// Add new corners, each iteration
		squareSide += 2

		// Calc min, max of new square
		minSquare = maxSquare + 1
		maxSquare += (squareSide - 1) * 4

		if square >= minSquare && square <= maxSquare {
			// Found inner square

			// Calc distante to first corner
			distanceToCorner := (squareSide - 2) - (square-minSquare)%(squareSide-1)
			if distanceToCorner > squareSide/2 {
				// Is closer to next corner than before corner
				distanceToCorner = (squareSide - distanceToCorner) - 1
			}
			circlePosition := squareSide / 2
			return circlePosition + (squareSide/2 - distanceToCorner)
		}

	}
	return -1
}

func repr(x, y int) string { return fmt.Sprintf("%d-%d", x, y) }

func CalcFirstLargerThan(threshold int) int {
	grid := map[string]int{}
	grid[repr(0, 0)] = 1
	x := 0
	y := 0
	currentDirection := [2]int{1, 0}
	perpendicularDirection := [2]int{0, 1}

	for grid[repr(x, y)] <= threshold {
		if _, exists := grid[repr(x+perpendicularDirection[0], y+perpendicularDirection[1])]; !exists {
			// Rotate counter clockwise when there is no filled square perpendicular to current direction
			currentDirection = perpendicularDirection
			perpendicularDirection[0] = currentDirection[1]
			perpendicularDirection[1] = -currentDirection[0]
		}
		x += currentDirection[0]
		y += currentDirection[1]

		grid[repr(x, y)] = 0
		// Check all 8 directions (sides and diagonals)
		for _, direction := range [][2]int{
			[2]int{1, 0},
			[2]int{1, 1},
			[2]int{0, 1},
			[2]int{-1, 1},
			[2]int{-1, 0},
			[2]int{-1, -1},
			[2]int{0, -1},
			[2]int{1, -1},
		} {
			value, exists := grid[repr(x+direction[0], y+direction[1])]
			if exists {
				grid[repr(x, y)] += value
			}

		}
	}
	return grid[repr(x, y)]

}

func main() {
	var inputSquare int
	fmt.Scanf("%d", &inputSquare)

	// Part 1
	result1 := CalcShortestPath(inputSquare)
	fmt.Println(result1)

	// Part 2
	result2 := CalcFirstLargerThan(inputSquare)
	fmt.Println(result2)
}
