package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

	if !hasPairRepetitions(s) {
		return false
	}

	if !hasAtLeastOneRepeatingLetterWithExactlyOneLetterBetweenThem(s) {
		return false
	}

	return true
}

func hasPairRepetitions(s string) bool {
	for i, letter := range s {
		j := i + 1

		if j == len(s) {
			continue
		}
		next := s[j]
		pair := string(letter) + string(next)

		if strings.Count(s, pair) > 1 {
			return true
		}
	}

	return false
}

func hasAtLeastOneRepeatingLetterWithExactlyOneLetterBetweenThem(s string) bool {

	for i, letter := range s {
		j := i + 2

		if j >= len(s) {
			continue
		}

		nextNext := s[j]

		if string(letter) == string(nextNext) {
			return true
		}
	}

	return false
}
