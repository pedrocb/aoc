package main

import (
	"encoding/json"
	"fmt"
)

func sum(input interface{}, filter string) int {
	switch inputType := input.(type) {
	case []interface{}:
		return sumArray(inputType, filter)
	case map[string]interface{}:
		return sumMap(inputType, filter)
	case float64:
		return int(inputType)
	}
	return 0
}

func sumMap(input map[string]interface{}, filter string) int {
	counter := 0

	for _, value := range input {
		if value == filter {
			return 0
		}
		counter += sum(value, filter)
	}
	return counter
}

func sumArray(input []interface{}, filter string) int {
	counter := 0
	for _, value := range input {
		counter += sum(value, filter)
	}
	return counter
}

func partOne(inputJson interface{}) int {
	counter := sum(inputJson, "")
	return counter
}

func partTwo(inputJson interface{}) int {
	counter := sum(inputJson, "red")
	return counter
}

func main() {
	var inputString string
	var inputJson interface{}
	fmt.Scanf("%s", &inputString)
	json.Unmarshal([]byte(inputString), &inputJson)
	fmt.Println(partOne(inputJson))
	fmt.Println(partTwo(inputJson))
}
