package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalln("Cannot open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var feet int64
	for scanner.Scan() {
		feet += howMuchRibbon(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Cannot scan", err)
	}

	fmt.Println(feet)

}

func howMuchRibbon(dimensions string) int64 {
	dArr := strings.Split(dimensions, "x")

	l, _ := strconv.ParseInt(dArr[0], 10, 64)
	w, _ := strconv.ParseInt(dArr[1], 10, 64)
	h, _ := strconv.ParseInt(dArr[2], 10, 64)

	min1, min2 := func(nums ...int64) (int64, int64) {
		min1 := int64(math.MaxInt64)
		min2 := int64(math.MaxInt64)

		for _, v := range nums {
			if v <= min1 {
				min2 = min1
				min1 = v
			}

			if v > min1 && v < min2 {
				min2 = v
			}
		}

		return min1, min2
	}(l, w, h)

	wrap := min1*2 + min2*2
	ribbon := l * w * h

	return wrap + ribbon
}
