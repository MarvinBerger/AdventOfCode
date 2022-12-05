package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./day1/input.txt")
	check(err)
	elves := make([]int, 1)
	currentElve := 0
	for _, element := range strings.Split(string(dat), "\r\n") {
		if len(element) > 0 {
			cal, _ := strconv.Atoi(element)
			elves[currentElve] += cal
		} else {
			currentElve++
			elves = append(elves, 0)
		}
	}

	sort.Ints(elves)
	println("Most Calories: " + strconv.Itoa(elves[len(elves)-1]))

	sumOfMost := 0
	for a := len(elves) - 3; a < len(elves); a++ {
		println(elves[a])
		sumOfMost += elves[a]
	}
	println("Top Three Sum: " + strconv.Itoa(sumOfMost))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
