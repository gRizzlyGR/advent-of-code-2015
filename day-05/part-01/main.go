package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var badPairs = []string{"ab", "cd", "pq", "xy"}
var vowels = []string{"a", "e", "i", "o", "u"}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		s := scanner.Text()

		if isNiceString(s) {
			count++
		}
	}

	fmt.Println(count)
}

func isNiceString(s string) bool {

	if contains(s, badPairs...) {
		return false
	}

	if !hasDoubleLetter(s) {
		return false
	}

	if !hasAtLeastThreeVowels(s) {
		return false
	}

	return true
}

func contains(s string, tokens ...string) bool {
	for _, token := range tokens {
		if strings.Contains(s, token) {
			return true
		}
	}

	return false
}

func hasDoubleLetter(s string) bool {
	for _, letter := range s {
		sLetter := string(letter)
		double := sLetter + sLetter
		if strings.Contains(s, double) {
			return true
		}
	}

	return false
}

func hasAtLeastThreeVowels(s string) bool {
	sliceContains := func(s string, coll ...string) bool {
		for _, elem := range coll {
			if string(elem) == s {
				return true
			}
		}

		return false
	}

	count := 0

	for _, letter := range s {
		if sliceContains(string(letter), vowels...) {
			count++
		}

		if count == 3 {
			return true
		}

	}

	return false
}
