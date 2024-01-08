package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	s := "test"
	fmt.Printf("%x", md5.Sum([]byte(s)))
}
