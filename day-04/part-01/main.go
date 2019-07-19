package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	key := "bgvyzdsv"

	i := 1
	for {
		data := []byte(fmt.Sprintf("%s%d", key, i))
		hash := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(hash, "00000") {
			fmt.Println(i, hash)
			break
		}
		i++
	}

}
