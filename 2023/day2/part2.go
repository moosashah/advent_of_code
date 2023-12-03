package main

import (
	"strconv"
	"strings"
)

func GameLoop(game string) (map[string]int, bool) {
	gameMax := map[string]int{
		"red":   1,
		"blue":  1,
		"green": 1,
	}
	_, rds := SplitGameAndRounds(game)
	if rds == nil {
		return nil, true
	}
	for _, v := range rds {
		mins := RoundLoop(v)
		for k, value := range mins {
			if gameMax[k] < value {
				gameMax[k] = value
			}
		}
	}
	return gameMax, false
}

func PowerOfSet(mins map[string]int) int {
	power := 1
	for _, value := range mins {
		power *= value
	}
	return power
}

func RoundLoop(r string) map[string]int {
	split := CleanString(strings.Split(r, ","))
	mins := map[string]int{
		"red":   1,
		"blue":  1,
		"green": 1,
	}
	for _, v := range split {
		sp := strings.Split(v, " ")
		colour := sp[1]
		quantity, _ := strconv.Atoi(sp[0])
		if mins[colour] < quantity {
			mins[colour] = quantity
		}
	}
	return mins
}

func Part2(input []string) int {
	var sum int
	for _, game := range input {
		max, err := GameLoop(game)
		if err {
			break
		}
		sum += PowerOfSet(max)
	}
	return sum
}
