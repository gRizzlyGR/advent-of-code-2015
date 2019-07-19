// x := uint16(123)
// y := uint16(456)
// d := x & y
// e := x | y
// f := x << 2
// g := y >> 2
// h := ^x
// i := ^y

// fmt.Printf("d: %v\ne: %v\nf: %v\ng: %v\nh: %v\ni: %v\nx: %v\ny: %v\n", d, e, f, g, h, i, x, y)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var varMap = make(map[string]uint16)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		fireCommand(s)
	}

	fmt.Println(varMap["a"])
	//fmt.Println(varMap)
}

func fireCommand(s string) {
	command := strings.Fields(s)
	switch len(command) {
	case 3:
		val, err := strconv.ParseUint(command[0], 10, 16) //strconv.Atoi(command[0])
		if err != nil {
			fmt.Println("Variable assignment found:", err)
			varMap[command[2]] = varMap[command[0]] // b -> a
		} else {
			varMap[command[2]] = uint16(val) // 123 -> a
		}
	case 4: // NOT a -> b
		varMap[command[3]] = ^varMap[command[1]]
	case 5:
		switch command[1] {
		case "AND": // a AND b -> c
			varMap[command[4]] = varMap[command[0]] & varMap[command[2]]
		case "OR": // a OR b -> c
			varMap[command[4]] = varMap[command[0]] | varMap[command[2]]
		case "LSHIFT": // a LSHIFT 1 -> b
			val, err := strconv.ParseUint(command[2], 10, 16)
			if err != nil {
				fmt.Println("LSHIFT error:", err)
			}
			varMap[command[4]] = varMap[command[0]] << uint16(val)
		case "RSHIFT": // a RSHIFT 1 -> b
			val, err := strconv.ParseUint(command[2], 10, 16)
			if err != nil {
				fmt.Println("RSHIFT error:", err)
			}
			varMap[command[4]] = varMap[command[0]] >> uint16(val)
		}
	}
}
