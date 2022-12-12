package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./day5/input.txt")
	stacks := make([][]string, 0)
	for _, line := range strings.Split(string(dat), "\r\n") {
		if line[0:3] == " 1 " {
			break
		}

		for i := 0; i <= len(line)/4; i++ {
			box := line[i*4 : i*4+3]
			stacks = addBoxToStack(i, box, stacks)
		}
	}

	for i := 0; i < len(stacks); i++ {
		ReverseSlice(stacks[i])
		stacks[i] = deleteEmptyBoxes(stacks[i])
	}

	for _, line := range strings.Split(string(dat), "\r\n") {
		if !strings.HasPrefix(line, "move") {
			continue
		}

		amount, src, dest := parseCommand(line)
		for i := 0; i < amount; i++ {
			stacks = moveBox(src, dest, stacks)
		}

	}

	boxes := ""
	for _, s := range stacks {
		boxes += s[len(s)-1]
	}

	boxes = strings.ReplaceAll(boxes, "[", "")
	boxes = strings.ReplaceAll(boxes, "]", "")
	fmt.Println(boxes)

}

func deleteEmptyBoxes(stack []string) []string {
	for i, e := range stack {
		if len(strings.Replace(e, " ", "", -1)) == 0 {
			return stack[0:i]
		}
	}
	return stack
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func moveBox(src int, dest int, stacks [][]string) [][]string {
	srcBox := stacks[src][len(stacks[src])-1]
	stacks[src] = stacks[src][0 : len(stacks[src])-1]
	stacks[dest] = append(stacks[dest], srcBox)

	return stacks
}

func addBoxToStack(stack int, box string, stacks [][]string) [][]string {
	if stack >= len(stacks) {
		stacks = append(stacks, make([]string, 0))
	}

	stacks[stack] = append(stacks[stack], box)
	return stacks
}

func parseCommand(line string) (int, int, int) {
	splitted := strings.Split(line, " ")
	fmt.Println(line)
	return stringToInt(splitted[1]), stringToInt(splitted[3]) - 1, stringToInt(splitted[5]) - 1
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
