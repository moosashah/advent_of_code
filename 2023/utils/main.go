package utils

import (
	"log"
	"os"
	"strings"
)

func Clean(input []string) []string {
	for i := range input {
		input[i] = strings.TrimSpace(input[i])
	}
	return input
}

func SplitInputIntoLines(input string) []string {
	lines := Clean(strings.Split(input, "\n"))
	lines = lines[:len(lines)-1]
	return lines
}

func ReadInput(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read from file: input.txt")
	}
	return string(content)
}
