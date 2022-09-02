package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction interface {
	apply([]byte) []byte
}

type Spin struct {
	size int
}

func (s Spin) apply(programs []byte) []byte {
	n := len(programs)
	result := make([]byte, n)
	for index, program := range programs {
		result[(index+s.size)%n] = program
	}
	return result
}

type Exchange struct {
	a, b int
}

func (e Exchange) apply(programs []byte) []byte {
	n := len(programs)
	result := make([]byte, n)

	copy(result, programs)
	result[e.a] = programs[e.b]
	result[e.b] = programs[e.a]
	return result

}

type Partner struct {
	a, b byte
}

func (p Partner) apply(programs []byte) []byte {
	n := len(programs)
	result := make([]byte, n)

	copy(result, programs)
	for index, program := range programs {
		if program == p.a {
			result[index] = p.b
		} else if program == p.b {
			result[index] = p.a
		}
	}
	return result

}

func ApplyInstructions(inputInstructions []instruction, n int) string {
	initialString := "abcdefghijklmnop"
	programs := []byte(initialString)
	for i := 0; i < n; i++ {
		for _, currentInstruction := range inputInstructions {
			programs = currentInstruction.apply(programs)
		}
		if string(programs) == initialString && i != 0 {
			cycleSize := i + 1
			numCycles := n / cycleSize
			i = (numCycles * cycleSize) - 1
		}
	}
	return string(programs)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawInstructions := strings.Split(scanner.Text(), ",")
	instructions := []instruction{}
	for _, rawInstruction := range rawInstructions {
		var inputInstruction instruction
		if rawInstruction[0] == 's' {
			size, _ := strconv.Atoi(rawInstruction[1:])
			inputInstruction = Spin{size: size}
		} else if rawInstruction[0] == 'x' {
			var a, b int
			fmt.Sscanf(rawInstruction, "x%d/%d", &a, &b)
			inputInstruction = Exchange{a: a, b: b}
		} else if rawInstruction[0] == 'p' {
			var a, b byte
			fmt.Sscanf(rawInstruction, "p%c/%c", &a, &b)
			inputInstruction = Partner{a: a, b: b}
		}
		instructions = append(instructions, inputInstruction)
	}

	// Part 1
	result1 := ApplyInstructions(instructions, 1)
	fmt.Println(result1)

	// Part 2
	result2 := ApplyInstructions(instructions, 1000000000)
	fmt.Println(result2)
}
