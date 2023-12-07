package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func part2(lines []string) string {
	scratchCards := make([]int, len(lines))
	for k := range lines {
		scratchCards[k] = 1
	}

	for i, card := range lines {
		sp := utils.Clean(strings.Split(card, " | "))
		winningNums := getWinningNums(sp[0])
		selectedNums := getSelectedNums(sp[1])
		if score := calculateWins(winningNums, selectedNums); score > 0 {
			for j := i + 1; j <= i+score; j++ {
				scratchCards[j] += scratchCards[i]
			}
		}
	}
	var sum int
	for _, v := range scratchCards {
		sum += v
	}
	return fmt.Sprint(sum)
}
