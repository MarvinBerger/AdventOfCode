package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./day3/input.txt")

	totalPriorityCommon := 0
	totalPriorityBadges := 0
	currentGroup := make([]string, 0)
	for _, backpack := range strings.Split(string(dat), "\r\n") {
		compartment2 := backpack[len(backpack)/2 : len(backpack)]
		currentGroup = append(currentGroup, backpack)
		foundCommonItem := false
		foundBadge := false
		for _, r := range backpack {
			if !foundCommonItem && strings.Contains(compartment2, string(r)) {
				totalPriorityCommon += toPriority(r)
				foundCommonItem = true
			}
			if !foundBadge && len(currentGroup) == 3 {
				if strings.Contains(currentGroup[0], string(r)) && strings.Contains(currentGroup[1], string(r)) {
					totalPriorityBadges += toPriority(r)
					currentGroup = make([]string, 0)
					foundBadge = true
				}
			}
			if foundBadge && foundCommonItem {
				break
			}
		}

	}
	fmt.Printf("CommonItemPriority: %d, BadgePriority: %d", totalPriorityCommon, totalPriorityBadges)
}

func toPriority(r rune) int {
	if r < 97 {
		return int(r - 38)
	} else {
		return int(r - 96)
	}
}
