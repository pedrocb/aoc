package main

import (
	"fmt"
	"math"
)

func partOne(inputNumber int) int {
	candidate := 1
	for {
		nPresents := 0
		divisors := []int{}
		for n := 1; n <= int(math.Sqrt(float64(candidate))+1); n++ {
			if candidate%n == 0 {
				divisors = append(divisors, n)
			}
		}
		for _, divisor := range divisors {
			nPresents += divisor
			if candidate != divisor*divisor {
				nPresents += candidate / divisor
			}
		}
		if nPresents*10 >= inputNumber {
			return candidate
		}
		candidate += 1
	}
}

func partTwo(inputNumber int) int {
	candidate := 1
	for {
		nPresents := 0
		divisors := []int{}
		for n := 1; n <= int(math.Sqrt(float64(candidate))+1); n++ {
			if candidate%n == 0 {
				divisors = append(divisors, n)
			}
		}
		for _, divisor := range divisors {
			if candidate/divisor <= 50 {
				nPresents += divisor
			}
			if candidate != divisor*divisor && divisor <= 50 {
				nPresents += candidate / divisor
			}
		}
		if nPresents*11 >= inputNumber {
			return candidate
		}
		candidate += 1
	}
}

func main() {
	var inputNumber int
	fmt.Scanf("%d", &inputNumber)
	fmt.Println(partOne(inputNumber))
	fmt.Println(partTwo(inputNumber))

}
