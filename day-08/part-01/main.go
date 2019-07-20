package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\\x([0-9]|[a-f]){2}`)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalln("Cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	charGap := 0
	for scanner.Scan() {
		s := scanner.Text()
		charGap += len(s) - len(unescape(s))
	}

	fmt.Println(charGap)
}

func unescape(s string) string {
	return parseHex(unescapeQuotes(unescapeSlashes(trimQuotes(s))))
}

func trimQuotes(s string) string {
	return s[1 : len(s)-1]
}

func unescapeSlashes(s string) string {
	return strings.ReplaceAll(s, `\\`, `\`)
}

func unescapeQuotes(s string) string {
	return strings.ReplaceAll(s, `\"`, `"`)
}

func parseHex(s string) string {
	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		s = strings.ReplaceAll(s, match, "@") // Any character is fine
	}

	return s
}
