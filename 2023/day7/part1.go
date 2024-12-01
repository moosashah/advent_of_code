package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

type player struct {
	cards    string
	bid      string
	rank     int
	handType int
}

const (
	five  = 6
	four  = 5
	full  = 4
	three = 3
	two   = 2
	one   = 1
	high  = 0
)

func part1(file string) {

	game := make([]player, 0)

	hands := utils.SplitInputIntoLines(utils.ReadInput(file))
	for _, hand := range hands {
		h := strings.Fields(hand)

		str := calculateStrength(h[0])
		game = append(game, player{cards: h[0], bid: h[1], handType: str})

	}

	for k, v := range game {
		fmt.Println(k, v)
	}

	sort.SliceStable(game, func(i, j int) bool {
		var str bool
		if game[i].handType == game[j].handType {
			str = calculateHighHand(game, i, j)
		} else {
			str = game[i].handType < game[j].handType
		}
		return str
	})

	for k, v := range game {
		v.rank = k + 1
		fmt.Println(v)
	}
}

func calculateHighHand(game []player, i, j int) bool {
	var str bool
	fmt.Println("calculate high hand")
	for key, value := range game[i].cards {

	}
	str = game[i].handType < game[j].handType
	return str
}

func calculateStrength(cards string) int {
	hand := make(map[string]int, 0)

	for _, card := range cards {
		hand[string(card)]++
	}

	counts := make([]int, 0)
	for _, v := range hand {
		counts = append(counts, v)
	}

	uniques := make(map[int]int, 0)
	for _, v := range counts {
		uniques[v]++
	}
	var strength int

	for k, v := range uniques {
		if k == 1 && v == 5 {
			strength = 0
		} else if k == 1 && v == 3 {
			strength = 1
		} else if k == 1 && v == 2 {
			strength = 3
		} else if k == 1 && v == 1 {
			if value, exists := uniques[4]; exists {
				if value == 1 {
					strength = 5
				}
			} else {
				strength = 2
			}
		} else if k == 2 && v == 2 {
			strength = 2
		} else if k == 3 && v == 1 {
			if value, exists := uniques[2]; exists {
				if value == 1 {
					strength = 4
				}
			} else {
				strength = 3
			}
		} else if k == 5 && v == 1 {
			strength = 6
		}
	}

	return strength
}
