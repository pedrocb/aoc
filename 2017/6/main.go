package main

import (
	"fmt"
)

func repr(banks []int) string {
	return fmt.Sprintf("%v", banks)
}

func redistribute(banks []int) {
	maxIndex := 0
	max := banks[0]
	for index := 1; index < len(banks); index++ {
		if banks[index] > max {
			max = banks[index]
			maxIndex = index
		}
	}
	banks[maxIndex] = 0
	for index := 0; index < len(banks); index++ {
		banks[index] += max / len(banks)
		mod := max % len(banks)
		relativePos := index - maxIndex - 1
		if relativePos < 0 {
			relativePos += len(banks)
		}
		if relativePos < mod {
			banks[index]++
		}
	}
}

func CountRedistributions(initialBanks []int) (int, int) {
	banks := make([]int, len(initialBanks))
	copy(banks, initialBanks)
	counter := 0
	cache := map[string]int{}
	exists := false
	for !exists {
		cache[repr(banks)] = counter
		counter++
		redistribute(banks)
		_, exists = cache[repr(banks)]
	}
	return counter, counter - cache[repr(banks)]
}

func main() {
	banks := []int{}
	for {
		var inputBank int
		_, err := fmt.Scanf("%d", &inputBank)
		if err != nil {
			break
		}
		banks = append(banks, inputBank)
	}

	// Part 1/2
	result1, result2 := CountRedistributions(banks)
	fmt.Println(result1)
	fmt.Println(result2)
}
