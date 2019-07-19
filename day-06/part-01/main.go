package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	row int
	col int
}

var arr = [1000][1000]int{}
var grid = arr[:]

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		light(s)
	}

	fmt.Println(howManyAreLit())

}

func light(s string) {
	command := strings.Split(s, " ")

	switch command[0] {
	case "toggle":
		toggle(getCoords(command[1]), getCoords(command[3]))
	case "turn":
		switch command[1] {
		case "on":
			turnOn(getCoords(command[2]), getCoords(command[4]))
		case "off":
			turnOff(getCoords(command[2]), getCoords(command[4]))
		}
	}
}

func getCoords(pair string) coords {
	a := strings.Split(pair, ",")
	row, _ := strconv.Atoi(a[0])
	col, _ := strconv.Atoi(a[1])
	return coords{row, col}
}

func toggle(from, to coords) {
	for i := from.row; i <= to.row; i++ {
		for j := from.col; j <= to.col; j++ {
			grid[i][j] = (grid[i][j] - 1) * -1
		}
	}
}

func turnOn(from, to coords) {
	for i := from.row; i <= to.row; i++ {
		for j := from.col; j <= to.col; j++ {
			grid[i][j] = 1
		}
	}
}

func turnOff(from, to coords) {
	for i := from.row; i <= to.row; i++ {
		for j := from.col; j <= to.col; j++ {
			grid[i][j] = 0
		}
	}
}

func howManyAreLit() int {
	count := 0
	for _, slice := range grid {
		for _, light := range slice {
			if light == 1 {
				count++
			}
		}
	}

	return count
}
