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
		feet += howMuchPaper(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Cannot scan", err)
	}

	fmt.Println(feet)

}

func howMuchPaper(dimensions string) int64 {
	dArr := strings.Split(dimensions, "x")

	var feet int64
	slack := int64(math.MaxInt64)

	// Multiply pairs of numbers
	arrSize := len(dArr)
	for i := 0; i < len(dArr); i++ {
		j := i + 1

		// The last one with the first one
		if j >= arrSize {
			j = 0
		}

		dim1, _ := strconv.ParseInt(dArr[i], 10, 64)
		dim2, _ := strconv.ParseInt(dArr[j], 10, 64)

		size := dim1 * dim2
		feet += size

		if size < slack {
			slack = size
		}
	}

	return 2*feet + slack
}
