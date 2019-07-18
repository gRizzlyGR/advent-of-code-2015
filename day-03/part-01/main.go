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

func main() {
	buffer, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}

	//buffer = []byte{'^', 'v', '^', 'v', '^', 'v', '^', 'v', '^', 'v'}
	//buffer = []byte{'^', '>', 'v', '<'}
	//buffer = []byte{'>'}
	fmt.Println(move(buffer))

}

func initCycle() map[byte][]byte {
	cycle := make(map[byte][]byte)

	cycle[n] = []byte{e, s, w}
	cycle[s] = []byte{w, n, e}
	cycle[e] = []byte{s, w, n}
	cycle[w] = []byte{n, e, s}

	return cycle
}

func initOpposite() map[byte]byte {
	opposite := make(map[byte]byte)

	opposite[n] = s
	opposite[s] = n
	opposite[e] = w
	opposite[w] = e

	return opposite
}

func move(buffer []byte) int {
	cycle := initCycle()
	opposite := initOpposite()

	sliceEqual := func(a, b []byte) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	count := 0
	for i, curr := range buffer {
		if i == 0 {
			count++
			continue
		}

		prev := buffer[i-1]

		if opposite[curr] == prev && (i-1 != 0) {
			fmt.Println(i, string(curr), string(prev))
			continue
		}

		if i < 3 {
			count++
			continue
		}

		prevPrev := buffer[i-2]
		prevPrevPrev := buffer[i-3]

		if sliceEqual(cycle[curr], []byte{prev, prevPrev, prevPrevPrev}) {
			fmt.Println(i, string(prevPrevPrev), string(prevPrev), string(prev), string(curr))
			continue
		}

		count++

	}

	return count
}
