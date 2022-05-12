package main

import (
	"fmt"
)

type reindeer struct {
	name       string
	velocity   int
	resistance int
	rest       int
}

func min(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func calcReindeerDistance(input reindeer, seconds int) int {
	cycle := input.resistance + input.rest
	return (seconds/cycle*input.resistance + min(input.resistance, seconds%cycle)) * input.velocity
}

func partOne(inputList []reindeer) int {
	maximum := 0
	for _, candidate := range inputList {
		candidateDistance := calcReindeerDistance(candidate, 2503)
		if candidateDistance > maximum {
			maximum = candidateDistance
		}
	}
	return maximum
}

func partTwo(inputList []reindeer) int {
	points := make(map[string]int)
	maxPoints := 0
	for second := 1; second <= 2503; second += 1 {
		maximum := 0
		maximumReindeer := []string{}
		for _, candidate := range inputList {
			distance := calcReindeerDistance(candidate, second)
			if distance > maximum {
				maximum = distance
				maximumReindeer = []string{candidate.name}
			} else if distance == maximum {
				maximumReindeer = append(maximumReindeer, candidate.name)
			}
		}
		for _, i := range maximumReindeer {
			points[i] += 1

			if points[i] > maxPoints {
				maxPoints = points[i]
			}
		}
	}
	return maxPoints
}

func main() {
	var inputList []reindeer
	for {
		var inputReindeer reindeer
		_, err := fmt.Scanf("%s can fly %d km/s for %d seconds, but then must rest for %d seconds.\n", &inputReindeer.name, &inputReindeer.velocity, &inputReindeer.resistance, &inputReindeer.rest)
		if err != nil {
			break

		}

		inputList = append(inputList, inputReindeer)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))

}
