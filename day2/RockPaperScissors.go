package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./day2/input.txt")
	totalScoreMethod1 := 0
	totalScoreMethod2 := 0
	for _, element := range strings.Split(string(dat), "\r\n") {
		col := strings.Split(element, " ")

		p1 := toString(col[0])
		p2 := toString(col[1])
		totalScoreMethod1 += toSymbolScore(col[1]) + match(p2, p1)

		totalScoreMethod2 += toSymbolScore(findCorrectSymbol(p1, toScore(col[1]))) + toScore(col[1])
	}
	fmt.Printf("Punktezahl M1: %d, M2: %d", totalScoreMethod1, totalScoreMethod2)
}

func findCorrectSymbol(p1 string, score int) string {
	symbols := []string{"SCHERE", "STEIN", "PAPIER"}
	for _, s := range symbols {
		if match(s, p1) == score {
			return s
		}
	}
	panic("Ergebnis konnte nicht gefunden werden")
}

func toSymbolScore(a string) int {
	switch a {
	case "A", "X", "STEIN":
		return 1
	case "B", "Y", "PAPIER":
		return 2
	case "C", "Z", "SCHERE":
		return 3
	}
	panic("Wert konnte nicht in Symbolscore umgerechnet werden")
}

func toScore(a string) int {
	switch a {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	panic("Wert konnte keiner Zielpunktezahl zugeordnet werden")
}

func toString(a string) string {
	switch a {
	case "A", "X":
		return "STEIN"
	case "B", "Y":
		return "PAPIER"
	case "C", "Z":
		return "SCHERE"
	}
	panic("Wert konnte keinem Symbol zugeordnet werden")
}

func match(u1 string, u2 string) int {

	both := []string{u1, u2}

	if u1 == u2 {
		return 3
	}

	if Contains(both, "SCHERE") && Contains(both, "STEIN") {
		if u1 == "SCHERE" {
			return 0
		} else {
			return 6
		}
	}

	if Contains(both, "STEIN") && Contains(both, "PAPIER") {
		if u1 == "STEIN" {
			return 0
		} else {
			return 6
		}
	}

	if Contains(both, "PAPIER") && Contains(both, "SCHERE") {
		if u1 == "PAPIER" {
			return 0
		} else {
			return 6
		}
	}

	panic("Matchup nicht möglich")
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
