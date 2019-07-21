package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var n = byte('^')
var s = byte('v')
var e = byte('>')
var w = byte('<')

type coords struct {
	row int
	col int
}

func main() {
	buffer, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}

	fmt.Println(move(buffer))

}

func move(path []byte) int {
	grid := make(map[coords]int)

	point := coords{0, 0}
	grid[point]++

	for _, dir := range path {
		point = next(point, dir)
		grid[point]++
	}

	return len(grid)
}

func next(point coords, direction byte) coords {
	switch direction {
	case n:
		return coords{point.row - 1, point.col}
	case s:
		return coords{point.row + 1, point.col}
	case e:
		return coords{point.row, point.col + 1}
	case w:
		return coords{point.row, point.col - 1}
	default:
		log.Fatalln("Unknown direction:", string(direction))
	}

	return coords{0, 0}
}
