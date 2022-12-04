package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	s := string(dat)
	fmt.Println(strings.Split(s, "\r\n"))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
