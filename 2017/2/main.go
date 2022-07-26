package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalcQuotient(row []int) int {
	numCols := len(row)
	for first := 0; first < numCols; first++ {
		for second := first + 1; second < numCols; second++ {
			var dividend, divisor int
			if row[first] > row[second] {
				dividend = row[first]
				divisor = row[second]
			} else {
				dividend = row[second]
				divisor = row[first]
			}
			if divisor == 0 {
				continue
			}
			if dividend%divisor == 0 {
				return dividend / divisor
			}

		}
	}
	return -1
}

func CalcMinMaxDiff(row []int) int {
	numCols := len(row)
	highest := row[0]
	lowest := row[0]
	for col := 1; col < numCols; col++ {
		if row[col] < lowest {
			lowest = row[col]
		}

		if row[col] > highest {
			highest = row[col]
		}
	}
	return highest - lowest
}

func CalcChecksum(spreadsheet [][]int, calcRowResult func([]int) int) int {
	checksum := 0
	numRows := len(spreadsheet)
	for row := 0; row < numRows; row++ {
		checksum += calcRowResult(spreadsheet[row])
	}

	return checksum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	spreadsheet := [][]int{}
	row := 0
	for scanner.Scan() {
		rowData := strings.Split(scanner.Text(), "\t")
		spreadsheet = append(spreadsheet, make([]int, len(rowData)))
		for id, digit := range rowData {
			number, _ := strconv.Atoi(digit)
			spreadsheet[row][id] = number
		}
		row++
	}
	// Part 1
	result1 := CalcChecksum(spreadsheet, CalcMinMaxDiff)
	fmt.Println(result1)

	// Part 2
	result2 := CalcChecksum(spreadsheet, CalcQuotient)
	fmt.Println(result2)
}
