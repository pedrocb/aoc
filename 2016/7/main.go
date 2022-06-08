package main

import (
	"fmt"
)

func supportTLS(ip string) bool {
	inBrackets := false
	foundABBA := false
	for currentCharIndex := 0; currentCharIndex <= len(ip)-4; currentCharIndex++ {
		if ip[currentCharIndex] == '[' {
			inBrackets = true
		} else if ip[currentCharIndex] == ']' {
			inBrackets = false
		} else if ip[currentCharIndex] == ip[currentCharIndex+3] && ip[currentCharIndex+1] == ip[currentCharIndex+2] && ip[currentCharIndex] != ip[currentCharIndex+1] {
			// ABBA is  found
			if inBrackets {
				return false
			} else {
				foundABBA = true
			}
		}
	}
	return foundABBA
}

func supportSSL(ip string) bool {
	foundABAs := map[string]bool{}
	foundBABs := map[string]bool{}

	inBrackets := false
	for currentCharIndex := 0; currentCharIndex <= len(ip)-3; currentCharIndex++ {
		if ip[currentCharIndex] == '[' {
			inBrackets = true
		} else if ip[currentCharIndex] == ']' {
			inBrackets = false
		} else if ip[currentCharIndex] == ip[currentCharIndex+2] && ip[currentCharIndex] != ip[currentCharIndex+1] && ip[currentCharIndex+1] != '[' && ip[currentCharIndex+1] != ']' {
			// ABA is  found

			if inBrackets {
				foundABAs[ip[currentCharIndex:currentCharIndex+2]] = true
				if foundBABs[ip[currentCharIndex+1:currentCharIndex+3]] {
					return true

				}
			} else {
				foundBABs[ip[currentCharIndex:currentCharIndex+2]] = true
				if foundABAs[ip[currentCharIndex+1:currentCharIndex+3]] {
					return true

				}
			}
		}
	}
	return false

}

func CountTLSSupportedIps(ips []string) int {
	counter := 0
	for _, ip := range ips {
		if supportTLS(ip) {
			counter++
		}
	}
	return counter
}

func CountSSLSupportedIps(ips []string) int {
	counter := 0
	for _, ip := range ips {
		if supportSSL(ip) {
			counter++
		}
	}
	return counter
}

func main() {
	var inputList []string
	for {
		var inputString string
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			break
		}

		inputList = append(inputList, inputString)
	}

	// Puzzle 1
	result1 := CountTLSSupportedIps(inputList)
	fmt.Println(result1)

	// Puzzle 2
	result2 := CountSSLSupportedIps(inputList)
	fmt.Println(result2)
}
