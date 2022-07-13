package main

import (
	"fmt"
	"strconv"
	"strings"
)

type operator interface {
	// Logic to scramble
	apply(string) string
	// Logic to unscramble
	revert(string) string
}

/*
   Swap operation
*/
type swap struct {
	first, second string
}

func (s swap) apply(input string) string {
	firstPosition, err := strconv.Atoi(s.first)
	if err != nil {
		// If not a number, find index of char
		firstPosition = strings.Index(input, s.first)
	}

	secondPosition, err := strconv.Atoi(s.second)
	if err != nil {
		// If not a number, find index of char
		secondPosition = strings.Index(input, s.second)
	}

	// Swap both positions
	result := []byte(input)
	result[firstPosition] = input[secondPosition]
	result[secondPosition] = input[firstPosition]
	return string(result)
}

func (s swap) revert(input string) string {
	return s.apply(input)
}

/*
   Rotate operation
*/

type rotate struct {
	offset string
}

func (r rotate) apply(input string) string {
	offset, err := strconv.Atoi(r.offset)
	if err != nil {
		// If not a number, find index
		offset = strings.Index(input, r.offset) + 1
		if offset > 4 {
			offset++
		}
	}
	// Module len(input) so we don't fully rotate multiple times
	offset = offset % len(input)

	if offset < 0 {
		// Convert left rotation to right rotation
		offset = len(input) + offset
	}
	return input[len(input)-offset:] + input[:len(input)-offset]
}

// How much is necessary to rotate if char is found on index i
// Calculated manually
var rotateCharPositions map[int]int = map[int]int{
	0: 7,
	1: -1,
	2: 2,
	3: -2,
	4: 1,
	5: -3,
	6: 0,
	7: -4,
}

func (r rotate) revert(input string) string {
	offset, err := strconv.Atoi(r.offset)
	if err != nil {
		// If not a number, check table to check how much is necessary to rotate
		charPosition := strings.Index(input, r.offset)
		offset = rotateCharPositions[charPosition]
	} else {
		// Otherwise, simple reverse it
		offset *= -1
	}
	return rotate{offset: strconv.Itoa(offset)}.apply(input)
}

/*
   Move operation
*/
type move struct {
	first, second int
}

func (m move) apply(input string) string {
	charToMove := input[m.first]
	result := input[:m.first] + input[m.first+1:]
	result = result[:m.second] + string(charToMove) + result[m.second:]
	return result
}

func (m move) revert(input string) string {
	return move{first: m.second, second: m.first}.apply(input)
}

/*
   Reverse operation
*/

type reverse struct {
	first, second int
}

func (r reverse) apply(input string) string {
	var reverseChunk strings.Builder
	for index := r.second; index >= r.first; index-- {
		reverseChunk.WriteByte(input[index])
	}
	return input[:r.first] + reverseChunk.String() + input[r.second+1:]
}

func (r reverse) revert(input string) string {
	return r.apply(input)
}

func Scramble(password string, operations []operator) string {
	for _, op := range operations {
		password = op.apply(password)
	}
	return password
}

func Unscramble(password string, operations []operator) string {
	for index := len(operations) - 1; index >= 0; index-- {
		password = operations[index].revert(password)
	}
	return password
}

func main() {
	operations := []operator{}

	for {
		var operation operator

		var operationName string
		_, err := fmt.Scanf("%s", &operationName)
		if err != nil {
			break
		}
		switch operationName {
		case "rotate":
			rotateOp := rotate{}

			var direction string
			fmt.Scanf("%s", &direction)
			if direction == "based" {
				fmt.Scanf("on position of letter %s\n", &rotateOp.offset)
			} else {
				fmt.Scanf("%s %s\n", &rotateOp.offset, new(string))
				if direction == "left" {
					rotateOp.offset = "-" + rotateOp.offset
				}

			}
			operation = rotateOp
		case "swap":
			swapOp := swap{}
			fmt.Scanf("%s %s with %s %s\n", new(string), &swapOp.first, new(string), &swapOp.second)
			operation = swapOp
		case "move":
			moveOp := move{}
			fmt.Scanf("position %d to position %d\n", &moveOp.first, &moveOp.second)
			operation = moveOp
		case "reverse":
			reverseOp := reverse{}
			fmt.Scanf("positions %d through %d\n", &reverseOp.first, &reverseOp.second)
			operation = reverseOp
		}

		operations = append(operations, operation)
	}

	// Part 1
	result1 := Scramble("abcdefgh", operations)
	fmt.Println(result1)

	// Part 2
	result2 := Unscramble("fbgdceah", operations)
	fmt.Println(result2)
}
