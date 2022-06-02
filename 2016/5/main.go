package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DecryptPasswordSimple(doorId string) string {
	var password strings.Builder
	for index := 0; password.Len() < 8; index++ {
		candidateInput := doorId + strconv.Itoa(index)
		candidateInputHashed := fmt.Sprintf("%x", md5.Sum([]byte(candidateInput)))

		if candidateInputHashed[:5] == "00000" {
			password.WriteByte(candidateInputHashed[5])
		}
	}
	return password.String()
}

func DecryptPasswordComplex(doorId string) string {
	password := [8]byte{}

	// Each cache i-th bit represents if the i-th byte is already found
	cache := 0
	// Once all 8 bits are filled, the password is found
	completeCache := int(math.Pow(2, 8)) - 1
	for index := 0; cache != completeCache; index++ {
		candidateInput := doorId + strconv.Itoa(index)
		candidateInputHashed := fmt.Sprintf("%x", md5.Sum([]byte(candidateInput)))

		if candidateInputHashed[:5] == "00000" {
			// Check if 6th byte is a number, less than 8 (size of password) and if that position's byte is not filled ()
			if position, err := strconv.Atoi(string(candidateInputHashed[5])); err == nil && position < 8 && cache>>position&1 == 0 {
				password[position] = candidateInputHashed[6]
				// Set position-th bit to 1
				cache |= 1 << position
			}
		}
	}
	return string(password[:])
}

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)

	// Puzzle 1
	result1 := DecryptPasswordSimple(inputString)
	fmt.Println(result1)

	// Puzzle 2
	result2 := DecryptPasswordComplex(inputString)
	fmt.Println(result2)
}
