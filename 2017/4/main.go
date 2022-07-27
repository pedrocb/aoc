package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func hasAnagrams(passphrase []string) bool {
	words := map[string]bool{}
	for _, word := range passphrase {
		orderedWord := []rune(word)
		sort.Slice(orderedWord, func(i, j int) bool { return orderedWord[i] < orderedWord[j] })
		if _, exists := words[string(orderedWord)]; exists {
			return false
		}
		words[string(orderedWord)] = true
	}
	return true
}

func hasDuplicates(passphrase []string) bool {
	words := map[string]bool{}
	for _, word := range passphrase {
		if _, exists := words[word]; exists {
			return false
		}
		words[word] = true
	}
	return true
}

func CountValidPassphrases(passphrases [][]string, isValidPassphrase func([]string) bool) int {
	counter := 0
	for _, passphrase := range passphrases {
		if isValidPassphrase(passphrase) {
			counter++
		}
	}
	return counter
}

func main() {
	passphrases := [][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		passphrases = append(passphrases, strings.Split(scanner.Text(), " "))
	}

	// Part 1
	result1 := CountValidPassphrases(passphrases, hasDuplicates)
	fmt.Println(result1)

	// Part 2
	result2 := CountValidPassphrases(passphrases, hasAnagrams)
	fmt.Println(result2)
}
