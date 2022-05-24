package main

import (
	"fmt"
)

func partOne(inputRow int, inputColumn int) int {
	row := 1
	column := 1
	currentValue := 20151125
	for row != inputRow || column != inputColumn {
		if row == 1 {
			row = column + 1
			column = 1
		} else {
			row -= 1
			column += 1
		}
		currentValue = (currentValue * 252533) % 33554393
	}
	return currentValue
}

func partTwo(inputRow int, inputColumn int) int {
	return -1
}

func main() {
	var inputRow, inputColumn int
	fmt.Scanf("To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &inputRow, &inputColumn)
	fmt.Println(partOne(inputRow, inputColumn))
	fmt.Println(partTwo(inputRow, inputColumn))

}
