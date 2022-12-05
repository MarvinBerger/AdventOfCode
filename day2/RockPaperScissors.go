package main

import (
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./day2/input.txt")
	for _, element := range strings.Split(string(dat), "\r\n") {
		round := strings.Split(element, " ")
		score, _ := toSSP(round[0])

	}
}

func toSSP(a string) (int, string) {
	switch a {
	case "A", "X":
		return 1, "SCHERE"
	case "B", "Y":
		return 2, "STEIN"
	case "C", "Z":
		return 3, "PAPIER"
	}
	return 0, ""
}

func compare(a string, b string) {
	_, u1 := toSSP(a)
	_, u2 := toSSP(b)

	both := []string{u1, u2}

	if Contains(both, "SCHERE") && Contains(both, "STEIN") {

	}

}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
