package main

import (
	"fmt"
	"github.com/pedrocb/aoc/common/assembunny"
)

func main() {
	inputInstructions := assembunny.ScanInstructions()
	// Part One
	result1 := assembunny.ExecInstructions(inputInstructions, map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	})["a"]
	fmt.Println(result1)

	// Part Two
	result2 := assembunny.ExecInstructions(inputInstructions, map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	})["a"]
	fmt.Println(result2)
}
