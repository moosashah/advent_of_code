package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	fmt.Printf("Part 1 example: %s\n", part1(utils.SplitInputIntoLines(exampleInput())))
	fmt.Printf("Part 1: %s\n", part1(utils.SplitInputIntoLines(utils.ReadInput("input.txt"))))
	fmt.Printf("Part 2: %s\n", part2(utils.SplitInputIntoLines(utils.ReadInput("input.txt"))))
}
