package main

import (
	"fmt"
)

func partOne(inputList []map[string]int) int {
	tickerTape := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	for indexSue, sue := range inputList {
		candidate := true
		for tickerTapeTrait, tickerTapeTraitValue := range tickerTape {
			sueTraitValue, exists := sue[tickerTapeTrait]
			if exists && sueTraitValue != tickerTapeTraitValue {
				candidate = false
				break
			}
		}
		if candidate {
			return indexSue
		}
	}
	return -1
}

func partTwo(inputList []map[string]int) int {
	tickerTape := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	for indexSue, sue := range inputList {
		candidate := true
		for tickerTapeTrait, tickerTapeTraitValue := range tickerTape {
			sueTraitValue, exists := sue[tickerTapeTrait]
			if exists {
				if tickerTapeTrait == "cats" || tickerTapeTrait == "trees" {
					if sueTraitValue <= tickerTapeTraitValue {
						candidate = false
						break
					}
				} else if tickerTapeTrait == "pomeranians" || tickerTapeTrait == "goldfish" {
					if sueTraitValue >= tickerTapeTraitValue {
						candidate = false
						break
					}

				} else {
					if sueTraitValue != tickerTapeTraitValue {
						candidate = false
						break
					}
				}
			}
		}
		if candidate {
			return indexSue
		}
	}
	return -1
}

func main() {
	var inputList []map[string]int
	for {
		inputTraits := make([]string, 3)
		inputTraitsValues := make([]int, 3)
		var index int
		_, err := fmt.Scanf("Sue %d: %s %d, %s %d, %s %d\n", &index,
			&inputTraits[0], &inputTraitsValues[0],
			&inputTraits[1], &inputTraitsValues[1],
			&inputTraits[2], &inputTraitsValues[2])
		if err != nil {
			break

		}
		inputSue := make(map[string]int)
		for index, inputTrait := range inputTraits {
			inputSue[inputTrait[:len(inputTrait)-1]] = inputTraitsValues[index]
		}
		inputList = append(inputList, inputSue)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))

}
