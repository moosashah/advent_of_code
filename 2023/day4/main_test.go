package main

import (
	"aoc/utils"
	"reflect"
	"testing"
)

func Test_getWinningNums(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Split",
			args: args{line: "Card 1: 41 48 83 86 17"},
			want: []int{41, 48, 83, 86, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinningNums(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWinningNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSelectedNums(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Selected",
			args: args{line: "83 86  6 31 17  9 48 53"},
			want: []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSelectedNums(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSelectedNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateScore(t *testing.T) {
	type args struct {
		winning  []int
		selected []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Score",
			args: args{
				winning:  []int{83, 86, 6, 31, 17, 9, 48, 53},
				selected: []int{41, 48, 83, 86, 17},
			},
			want: 4,
		},
		{
			name: "No points",
			args: args{
				winning:  []int{1, 2, 3, 4},
				selected: []int{41, 48, 83, 86, 17},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWins(tt.args.winning, tt.args.selected); got != tt.want {
				t.Errorf("calculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Part 1 with example input",
			args: args{input: utils.SplitInputIntoLines(exampleInput())},
			want: "13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Part 2 with example input",
			args: args{lines: utils.SplitInputIntoLines(exampleInput())},
			want: "29",
		},
		{
			name: "Part 2",
			args: args{lines: utils.SplitInputIntoLines(utils.ReadInput("input.txt"))},
			want: "8570000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
