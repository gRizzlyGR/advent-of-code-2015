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

	for _, b := range buffer {
		if b == leftP {
			floors++
		} else if b == rightP {
			floors--
		} else {
			log.Fatalln("Unrecognized symbol:", string(b))
		}
	}

	fmt.Println(floors)

}
