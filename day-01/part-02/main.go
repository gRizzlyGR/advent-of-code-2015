package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	buffer, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatalln("File cannot be opened", err)
	}

	leftP := byte('(')
	rightP := byte(')')

	floors := 0

	for i, b := range buffer {
		if b == leftP {
			floors++
		} else if b == rightP {
			floors--
		} else {
			log.Fatalln("Unrecognized symbol:", string(b))
		}

		if floors == -1 {
			fmt.Println("Position for basement:", i+1)
			break
		}
	}
}
