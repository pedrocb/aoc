package main

import (
	"fmt"
)

const (
	Rect      = "rect"
	RotateRow = "row"
	RotateCol = "column"
)

var screenWidth = 50
var screenHeight = 6

type instruction struct {
	command string
	x, y    int
	offset  int
}

func CountLitPixels(inputInstructions []instruction) ([]bool, int) {
	counter := 0
	grid := make([]bool, screenHeight*screenWidth)
	for _, currentInstruction := range inputInstructions {
		switch currentInstruction.command {
		case Rect:
			for x := 0; x < currentInstruction.x; x++ {
				for y := 0; y < currentInstruction.y; y++ {
					if !grid[x+screenWidth*y] {
						// Only increase if the led was light up
						counter += 1
						grid[x+screenWidth*y] = true
					}
				}
			}
		case RotateRow:
			// Create copy of previous row, so it is not changed
			previousRow := make([]bool, screenWidth)
			copy(previousRow, grid[currentInstruction.y*screenWidth:screenWidth+currentInstruction.y*screenWidth])
			for x := 0; x < screenWidth; x++ {
				offsetX := x + currentInstruction.offset
				if offsetX >= screenWidth {
					// It overflows
					offsetX = offsetX % screenWidth
				}
				grid[offsetX+currentInstruction.y*screenWidth] = previousRow[x]
			}
		case RotateCol:
			// Create copy of previous col, so it is not changed
			previousCol := make([]bool, screenHeight)
			for y := 0; y < screenHeight; y++ {
				previousCol[y] = grid[currentInstruction.x+y*screenWidth]
			}
			for y := 0; y < screenHeight; y++ {
				offsetY := y + currentInstruction.offset
				if offsetY >= screenHeight {
					// It overflows
					offsetY = offsetY % screenHeight
				}
				grid[currentInstruction.x+offsetY*screenWidth] = previousCol[y]
			}
		}
	}

	return grid, counter
}

func main() {
	var inputList []instruction
	for {
		var inputInstruction instruction
		_, err := fmt.Scanf("%s", &inputInstruction.command)
		if err != nil {
			break
		}
		switch inputInstruction.command {
		case "rect":
			fmt.Scanf("%dx%d", &inputInstruction.x, &inputInstruction.y)
		case "rotate":
			fmt.Scanf("%s", &inputInstruction.command)
			if inputInstruction.command == RotateRow {
				fmt.Scanf("y=%d by %d", &inputInstruction.y, &inputInstruction.offset)
			} else {
				fmt.Scanf("x=%d by %d", &inputInstruction.x, &inputInstruction.offset)
			}
		}

		inputList = append(inputList, inputInstruction)
	}

	result2, result1 := CountLitPixels(inputList)
	fmt.Println(result1)

	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			if result2[x+y*screenWidth] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
