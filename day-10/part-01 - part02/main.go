package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide a valid number as argument!")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Please provide a valid number:", err)
	}

	s := "1113222113"

	fmt.Println(os.Args[1])
	for i := 0; i < n; i++ {
		fmt.Println(i)
		s = look(s)
	}
	fmt.Println(len(s))
}

func look(s string) string {
	count := 0
	var prev rune
	say := ""
	for i, c := range s {
		if i == 0 {
			prev = c
			count++
			continue
		}

		if c == prev {
			count++
		} else {
			say += fmt.Sprintf("%d%s", count, string(prev))
			count = 1
		}
		prev = c
	}
	say += fmt.Sprintf("%d%s", count, string(prev))
	return say
}
