package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./day6/input.txt")
	currentString := make([]rune, 0)
	for index, char := range string(dat) {
		currentString = append(currentString, char)

		if len(currentString) >= 4 && uniqueArray(currentString[len(currentString)-4:]) {
			fmt.Printf("Found Unique 4! %d %s \r\n", index+1, string(currentString[len(currentString)-4:]))
		}

		if len(currentString) >= 14 && uniqueArray(currentString[len(currentString)-14:]) {
			fmt.Printf("Found Unique 14! %d %s \r\n", index+1, string(currentString[len(currentString)-14:]))
		}
	}
}

func uniqueArray(chars []rune) bool{
	for index, element := range chars {
		if !(findFirstIndex(chars, element) == index) {
			return false
		}
	}
	return true
}

func findFirstIndex(chars []rune, r rune) int {
	for index, element := range chars {
		if element == r {
			return index
		}
	}
	return -1
}