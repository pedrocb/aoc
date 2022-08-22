package main

import (
	"fmt"
	"github.com/pedrocb/aoc/common/knothash"
	"strconv"
)

func genGrid(key string) map[int]bool {
	grid := make(map[int]bool, 128*128)
	for y := 0; y < 128; y++ {
		rowKey := fmt.Sprintf("%s-%d", key, y)
		hash := knothash.KnotHash(rowKey)
		// For each hexadecimal character
		for charPos, char := range hash {
			i, _ := strconv.ParseInt(string(char), 16, 0)
			// For each bit in the hex char
			for bit := 0; bit < 4; bit++ {
				x := charPos*4 + bit
				grid[x+128*y] = (i>>(3-bit)&1 == 1)
			}
		}

	}
	return grid
}

func CountSquares(grid map[int]bool) int {
	counter := 0
	for i := 0; i < 128*128; i++ {
		if grid[i] {
			counter++
		}
	}
	return counter
}

func CountGroups(grid map[int]bool) int {
	groups := map[int]bool{}
	nGroups := 0
	directions := [][2]int{[2]int{-1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}}
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if !groups[x+y*128] && grid[x+y*128] {
				nGroups++
				queue := [][2]int{[2]int{x, y}}
				for len(queue) > 0 {
					currentX := queue[0][0]
					currentY := queue[0][1]
					queue = queue[1:]
					if currentX < 0 || currentY < 0 || currentX >= 128 || currentY >= 128 {
						continue
					}
					currentPos := currentX + currentY*128
					if !groups[currentPos] && grid[currentPos] {
						groups[currentPos] = true
						for _, dir := range directions {
							queue = append(queue, [2]int{currentX + dir[0], currentY + dir[1]})
						}
					}
				}
			}

		}
	}
	return nGroups
}

func main() {
	var key string
	fmt.Scanf("%s", &key)

	grid := genGrid(key)

	// Part 1
	result1 := CountSquares(grid)
	fmt.Println(result1)

	// Part 2
	result2 := CountGroups(grid)
	fmt.Println(result2)
}
