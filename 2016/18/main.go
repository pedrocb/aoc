package main

import (
	"fmt"
)

// Possible cell values
const (
	SAFE = false
	TRAP = true
)

func countSafeTiles(firstRow []bool, n int) int {
	rowSize := len(firstRow)
	grid := [][]bool{firstRow}

	// Counter of safe tiles
	counter := 0
	for currentRow := 0; currentRow < n; currentRow++ {
		for currentCol := 0; currentCol < rowSize; currentCol++ {
			if currentRow == 0 {
				// If first position. pick from the already filled grid (from input)
				if grid[currentRow][currentCol] == SAFE {
					counter++
				}
			} else {
				var left, right bool
				if currentCol == 0 {
					// If first col, initialize slice of currentRow
					grid = append(grid, []bool{})

					// If first col, left is safe since it is a wall
					left = SAFE
				} else {
					left = grid[currentRow-1][currentCol-1]
				}

				if currentCol == rowSize-1 {
					// If last col, right is safe since it is a wall
					right = SAFE
				} else {
					right = grid[currentRow-1][currentCol+1]
				}

				center := grid[currentRow-1][currentCol]

				if (left == TRAP && center == TRAP && right == SAFE) ||
					(left == SAFE && center == TRAP && right == TRAP) ||
					(left == TRAP && center == SAFE && right == SAFE) ||
					(left == SAFE && center == SAFE && right == TRAP) {
					grid[currentRow] = append(grid[currentRow], TRAP)
				} else {
					grid[currentRow] = append(grid[currentRow], SAFE)
					counter++
				}
			}
		}
	}
	return counter
}

func main() {
	// Represent each cell as bool
	firstRow := []bool{}
	for {
		var inputChar byte
		_, err := fmt.Scanf("%c", &inputChar)

		if err != nil {
			break
		}

		if inputChar == '.' {
			firstRow = append(firstRow, SAFE)
		} else if inputChar == '^' {
			firstRow = append(firstRow, TRAP)
		}
	}

	// Part 1
	result1 := countSafeTiles(firstRow, 40)
	fmt.Println(result1)

	// Part 2
	result2 := countSafeTiles(firstRow, 400000)
	fmt.Println(result2)
}
