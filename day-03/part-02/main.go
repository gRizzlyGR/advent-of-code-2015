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

	// The two Santas start delivering at the same house.
	// It's not needed to increment by two,actually
	// but this way it's compliant to the story
	origin := coords{0, 0}
	grid[origin] += 2

	c := make(chan coords)
	go deliver(path, 0, 2, origin, c) // Santa
	go deliver(path, 1, 2, origin, c) // Robo-Santa

	updateMap(grid, c, len(path))

	close(c)

	return len(grid)
}

func deliver(path []byte, start int, step int, point coords, c chan<- coords) {
	for i := start; i < len(path); i += step {
		point = next(point, path[i])
		c <- point
	}
}

func updateMap(grid map[coords]int, c <-chan coords, limit int) {
	for i := 0; i < limit; i++ {
		grid[<-c]++
	}
}

func next(origin coords, direction byte) coords {
	switch direction {
	case n:
		return coords{origin.row - 1, origin.col}
	case s:
		return coords{origin.row + 1, origin.col}
	case e:
		return coords{origin.row, origin.col + 1}
	case w:
		return coords{origin.row, origin.col - 1}
	default:
		log.Fatalln("Unknown direction:", string(direction))
	}

	return coords{0, 0}
}
