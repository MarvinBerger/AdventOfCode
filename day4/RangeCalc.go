package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./day4/input.txt")
	pairsInRange := 0
	overlap := 0
	for _, line := range strings.Split(string(dat), "\r\n") {
		r1, r2 := splitLine(line)
		if rangeInRange(r1, r2) || rangeInRange(r2, r1) {
			pairsInRange++
		}
		if rangeOverlapsABit(r1, r2) || rangeOverlapsABit(r2, r1) {
			overlap++
		}
	}
	fmt.Printf("RedundantPairs: %d, Overlap: %d", pairsInRange, overlap)
}

func splitLine(line string) (string, string) {
	ranges := strings.Split(line, ",")
	return ranges[0], ranges[1]
}

func splitRange(r string) (int, int) {
	numbers := strings.Split(r, "-")
	n1, _ := strconv.Atoi(numbers[0])
	n2, _ := strconv.Atoi(numbers[1])
	return n1, n2
}

func rangeInRange(r1 string, r2 string) bool {
	r1b, r1e := splitRange(r1)
	r2b, r2e := splitRange(r2)

	return r1b >= r2b && r1e <= r2e
}

func rangeOverlapsABit(r1 string, r2 string) bool {
	r1b, _ := splitRange(r1)
	r2b, r2e := splitRange(r2)

	return r1b >= r2b && r1b <= r2e
}