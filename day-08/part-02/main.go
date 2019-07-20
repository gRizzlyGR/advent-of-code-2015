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
		es := escape(s)
		charGap += len(es) - len(s)
		fmt.Println(len(s), s)
		fmt.Println(len(es), es)
		fmt.Println()
	}

	fmt.Println(charGap)

	x := []string{`""`, `"abc"`, `"aaa\"aaa"`, `"\x27"`}

	a := 0
	for _, v := range x {
		y := escape(v)
		fmt.Println(len(v), v)
		fmt.Println(len(y), y)
		fmt.Println()
		a += len(y) - len(v)
	}
	fmt.Println(a)

}

func unescape(s string) string {
	return parseHex(unescapeQuotes(unescapeSlashes(trimQuotes(s))))
}

func escape(s string) string {
	return escapeSurroundingQuotes(escapeEscapedQuotes(escapeHex(s)))
}

func trimQuotes(s string) string {
	return s[1 : len(s)-1]
}

func unescapeSlashes(s string) string {
	return strings.ReplaceAll(s, `\\`, `\`)
}

func escapeEscapedSlashes(s string) string {
	return strings.ReplaceAll(s, `\\`, `\\\\`)
}

func unescapeQuotes(s string) string {
	return strings.ReplaceAll(s, `\"`, `"`)
}

func escapeSurroundingQuotes(s string) string {
	s = `"\` + s
	s = s[:len(s)-1] + `\""`

	return s
}

func escapeEscapedQuotes(s string) string {
	if strings.HasSuffix(s, `\"`) {
		fmt.Println("+++++++++++++++++++", s)
		s = strings.ReplaceAll(s[:len(s)-1], `\"`, `\\\"`)
		s = s + `"`
		fmt.Println("--------------------", s)
		return s
	}
	return strings.ReplaceAll(s, `\"`, `\\\"`)
}

func escapeEscQuotes(s string) string {
	return ""
}

func parseHex(s string) string {
	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		s = strings.ReplaceAll(s, match, "@") // Any character is fine
	}

	return s
}

func escapeHex(s string) string {
	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		s = strings.ReplaceAll(s, match, `\`+match)
	}

	return s
}
