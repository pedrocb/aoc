package main

import (
	"fmt"
	"strconv"
)

func SolveCaptcha(captcha string, interval int) int {
	solution := 0

	for index, digit := range captcha {
		if byte(digit) == captcha[(index+interval)%len(captcha)] {
			digitInt, _ := strconv.Atoi(string(digit))
			solution += digitInt
		}
	}
	return solution
}

func main() {
	var captcha string
	fmt.Scanf("%s", &captcha)

	// Part 1
	result1 := SolveCaptcha(captcha, 1)
	fmt.Println(result1)

	// Part 2
	result2 := SolveCaptcha(captcha, len(captcha)/2)
	fmt.Println(result2)
}
