package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convertTextToNums(line string) string {
	numsDict := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for k, v := range numsDict {
		safeInsertIdx := 2
		rep := k[:safeInsertIdx] + v + k[safeInsertIdx:]
		line = strings.ReplaceAll(line, k, rep)
	}
	return line
}

func findNumbers(line string) (int, int, error) {
	numRegex := regexp.MustCompile(`\d`)
	xs := numRegex.FindAllString(line, -1)
	if len(xs) == 0 {
		return 0, 0, fmt.Errorf("no numbers found in line")
	}

	startNum, err := strconv.Atoi(xs[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to parse start number: %#w", err)
	}
	endNum, err := strconv.Atoi(xs[len(xs)-1])
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to parse end number: %#w", err)
	}
	return startNum, endNum, nil
}

func concatNumbers(startNum, endNum int) (int, error) {
	numStr := fmt.Sprintf("%d%d", startNum, endNum)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse joined number: %w", err)
	}
	return num, nil
}

func sumAllNumbers(nums []int) int {
	var n int
	for _, num := range nums {
		n += num
	}
	return n
}

func everything(file string) int {
	content, err := os.Open(file)
	if err != nil {
		log.Fatalf("Could not read from file: %s", file)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)

	nums := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		l := convertTextToNums(line)
		fmt.Printf("Converted %s to %s\n", line, l)
		startNum, endNum, err := findNumbers(l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("start num: %d, end num: %d\n", startNum, endNum)
		num, err := concatNumbers(startNum, endNum)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Joined number: %d\n", num)
		nums = append(nums, num)
	}
	return sumAllNumbers(nums)
}
