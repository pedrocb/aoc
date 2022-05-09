package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func partOne(inputString string) int {
	for i := 0; i < 999999; i++ {
		candidate := inputString + strconv.Itoa(i)
		result := md5.Sum([]byte(candidate))
		hexResult := fmt.Sprintf("%x", result)
		if hexResult[:5] == "00000" {
			return i
		}
	}
	return -1

}

func partTwo(inputString string) int {
	for i := 0; i < 9999999; i++ {
		candidate := inputString + strconv.Itoa(i)
		result := md5.Sum([]byte(candidate))
		hexResult := fmt.Sprintf("%x", result)
		if hexResult[:6] == "000000" {
			return i
		}
	}
	return -1

}

func main() {
	var inputString string
	for {
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}
	}
	fmt.Println(partOne(inputString))
	fmt.Println(partTwo(inputString))
}
