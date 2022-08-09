package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directionVectors map[string][3]int = map[string][3]int{
	"n":  [3]int{0, 1, -1},
	"s":  [3]int{0, -1, 1},
	"ne": [3]int{1, 0, -1},
	"sw": [3]int{-1, 0, 1},
	"nw": [3]int{-1, 1, 0},
	"se": [3]int{1, -1, 0},
}

func CalcSteps(directions []string) (int, int) {
	posX := 0
	posY := 0
	posZ := 0
	maxSteps := 0
	xDistance := 0
	yDistance := 0
	zDistance := 0
	for _, dir := range directions {
		posX += directionVectors[dir][0]
		posY += directionVectors[dir][1]
		posZ += directionVectors[dir][2]

		xDistance = posX
		if xDistance < 0 {
			xDistance *= -1
		}
		yDistance = posY
		if yDistance < 0 {
			yDistance *= -1
		}
		zDistance = posZ
		if zDistance < 0 {
			zDistance *= -1
		}
		if (xDistance+yDistance+zDistance)/2 > maxSteps {
			maxSteps = (xDistance + yDistance + zDistance) / 2
		}

	}
	return (xDistance + yDistance + zDistance) / 2, maxSteps

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	directions := strings.Split(scanner.Text(), ",")

	// Part 1
	result1, result2 := CalcSteps(directions)
	fmt.Println(result1)
	fmt.Println(result2)
}
