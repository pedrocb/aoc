package main

import (
	"fmt"
	"sort"
)

type instruction struct {
	lowOutput, highOutput string

	bot   string
	value int
}

func buildState(inputInstructions []instruction) map[string][]int {
	instructionsMap := map[string][]instruction{}
	outputsList := []string{}

	// Create map to easily access instructions that produce the values of a specific bot
	for _, currentInstruction := range inputInstructions {
		if currentInstruction.value == 0 {
			instructionsMap[currentInstruction.lowOutput] = append(instructionsMap[currentInstruction.lowOutput], currentInstruction)
			instructionsMap[currentInstruction.highOutput] = append(instructionsMap[currentInstruction.highOutput], currentInstruction)
			if currentInstruction.lowOutput[:3] != "bot" {
				outputsList = append(outputsList, currentInstruction.lowOutput)
			}
			if currentInstruction.highOutput[:3] != "bot" {
				outputsList = append(outputsList, currentInstruction.highOutput)
			}
		} else {
			instructionsMap[currentInstruction.bot] = append(instructionsMap[currentInstruction.bot], currentInstruction)
		}
	}
	state := map[string][]int{}
	for _, output := range outputsList {
		// Build map with each bot as key and its values as value, by recursively applying instructions for the outputs
		calc(output, state, instructionsMap)
	}
	return state
}

func calc(bot string, state map[string][]int, instructionsMap map[string][]instruction) {
	for _, currentInstruction := range instructionsMap[bot] {
		if currentInstruction.value == 0 {
			// The value comes from another bot, so both values of that bot must be calculated, so we can know if which is the low or high value
			if _, exists := state[currentInstruction.bot]; !exists {
				calc(currentInstruction.bot, state, instructionsMap)
			}

			// Append value to the bot accordingly (i.e if low or high value)
			if bot == currentInstruction.lowOutput {
				state[bot] = append(state[bot], state[currentInstruction.bot][0])
			} else {
				state[bot] = append(state[bot], state[currentInstruction.bot][1])
			}
		} else {
			// Value instruction i.e the value is simply added to a bot
			state[bot] = append(state[bot], currentInstruction.value)
		}
		sort.Ints(state[bot])
	}
}

func FindBotThatCompares(low int, high int, state map[string][]int) string {
	for bot, value := range state {
		if len(value) == 2 && value[0] == low && value[1] == high {
			return bot
		}
	}
	return ""
}

func main() {
	var inputInstructions []instruction
	for {
		var inputString string
		var inputInstruction instruction
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
		if inputString == "value" {
			fmt.Scanf("%d goes to bot %s", &inputInstruction.value, &inputInstruction.bot)
			inputInstruction.bot = fmt.Sprintf("bot-%s", inputInstruction.bot)
		} else {
			var lowOutputType, lowOutputId, highOutputType, highOutputId string
			fmt.Scanf("%s gives low to %s %s and high to %s %s ", &inputInstruction.bot, &lowOutputType, &lowOutputId, &highOutputType, &highOutputId)
			inputInstruction.bot = fmt.Sprintf("%s-%s", inputString, inputInstruction.bot)
			inputInstruction.lowOutput = fmt.Sprintf("%s-%s", lowOutputType, lowOutputId)
			inputInstruction.highOutput = fmt.Sprintf("%s-%s", highOutputType, highOutputId)
		}

		inputInstructions = append(inputInstructions, inputInstruction)
	}
	state := buildState(inputInstructions)

	result1 := FindBotThatCompares(17, 61, state)
	fmt.Println(result1)

	result2 := state["output-0"][0] * state["output-1"][0] * state["output-2"][0]
	fmt.Println(result2)
}
