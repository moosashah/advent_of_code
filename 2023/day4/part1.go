package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func exampleInput() string {
	return `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
}

func getWinningNums(line string) []int {
	numsSlice := make([]int, 0)
	winningNums := utils.Clean(strings.Split(line, ":"))[1]
	for _, v := range utils.Clean(strings.Split(winningNums, " ")) {
		n, _ := strconv.Atoi(v)
		if n == 0 {
			continue
		}
		numsSlice = append(numsSlice, n)
	}
	return numsSlice
}

func getSelectedNums(line string) []int {
	numsSlice := make([]int, 0)
	for _, v := range utils.Clean(strings.Split(line, " ")) {
		n, _ := strconv.Atoi(v)
		if n == 0 {
			continue
		}
		numsSlice = append(numsSlice, n)
	}
	return numsSlice
}

func calculateWins(winning, selected []int) int {
	count := 0
	for _, w := range winning {
		for _, s := range selected {
			if w == s {
				count++
				break
			}
		}
	}
	return count
}

func part1(lines []string) string {
	var sum int
	for _, line := range lines {
		sp := utils.Clean(strings.Split(line, " | "))
		winningNums := getWinningNums(sp[0])
		selectedNums := getSelectedNums(sp[1])
		sum += int(math.Pow(2, float64(calculateWins(winningNums, selectedNums)-1)))
	}
	return fmt.Sprint(sum)
}
