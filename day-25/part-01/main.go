package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	row, _ := strconv.Atoi(os.Args[1])
	col, _ := strconv.Atoi(os.Args[2])
	i, j := 1, 1
	val := 20151125
	for i < max(row, col)*2 {
		endI, endJ := j, i
		for i != endI && j != endJ {
			i--
			j++
			val = (val * 252533) % 33554393
			if i == row && j == col {
				fmt.Println(val)
				return
			}
		}
		i = endJ
		j = 1
		val = (val * 252533) % 33554393
		if i == row && j == col {
			fmt.Println(val)
			return
		}
		i++
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
