package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadExample() string {
	return `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
}

func Limits() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func ReadContent(input string) []string {
	s := strings.Split(input, "\n")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}

func GetGameNumber(g string) int {
	nr := regexp.MustCompile(`\d+`)
	gameNr := nr.FindAllString(g, -1)
	v, _ := strconv.Atoi(gameNr[0])
	return v
}

func playExample() int {
	var sum int
	games := ReadContent(ReadExample())
	for _, game := range games {
		sum += IsGamePossible(game)
	}
	return sum
}

func SplitRounds(rounds string) []string {
	return CleanString(strings.Split(rounds, ";"))
}

func SplitGameAndRounds(game string) (int, []string) {
	g := CleanString(strings.Split(game, ":"))
	gameHeading := g[0]
	if gameHeading == "" {
		return 0, nil
	}
	return GetGameNumber(gameHeading), SplitRounds(g[1])
}

func IsGamePossible(game string) int {
	gameNr, rds := SplitGameAndRounds(game)
	if rds == nil {
		return 0
	}
	for _, v := range rds {
		possible := IsRoundPossible(v, Limits())
		if !possible {
			return 0
		}
	}
	return gameNr
}

func CleanString(s []string) []string {
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}

func IsRoundPossible(r string, limits map[string]int) bool {
	split := CleanString(strings.Split(r, ","))
	p := true
	for _, v := range split {
		sp := strings.Split(v, " ")
		colour := sp[1]
		quantity, _ := strconv.Atoi(sp[0])
		if limits[colour] < quantity {
			p = false
			break
		}
	}
	return p
}

func ReadInput() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not read from file: input.txt")
	}
	return string(content)
}

func part1() int {
	var sum int
	games := ReadContent(ReadInput())
	for _, game := range games {
		gameNr := IsGamePossible(game)
		sum += gameNr
	}
	return sum
}
