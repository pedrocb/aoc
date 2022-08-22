package main

import (
	"bufio"
	"fmt"
	"github.com/pedrocb/aoc/common/knothash"
	"os"
	"strconv"
	"strings"
)

func main() {
	lengths := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()
	for _, inputNumber := range strings.Split(inputString, ",") {
		number, _ := strconv.Atoi(inputNumber)
		lengths = append(lengths, number)
	}

	// Part 1
	result1 := knothash.KnotHashInternals(lengths, 1)
	fmt.Println(result1[0] * result1[1])

	// Part 2
	result2 := knothash.KnotHash(inputString)
	fmt.Println(result2)

}
