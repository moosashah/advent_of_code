package main

import (
	"reflect"
	"testing"
)

func TestPowerOfSet(t *testing.T) {
	type args struct {
		mins map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Game 1",
			args: args{mins: map[string]int{
				"red":   4,
				"green": 2,
				"blue":  6,
			}},
			want: 48,
		},
		{
			name: "Game 2",
			args: args{mins: map[string]int{
				"red":   1,
				"green": 3,
				"blue":  4,
			}},
			want: 12,
		},
		{
			name: "Game 3",
			args: args{mins: map[string]int{
				"red":   20,
				"green": 13,
				"blue":  6,
			}},
			want: 1560,
		},
		{
			name: "Game 4",
			args: args{mins: map[string]int{
				"red":   14,
				"green": 3,
				"blue":  15,
			}},
			want: 630,
		},
		{
			name: "Game 5",
			args: args{mins: map[string]int{
				"red":   6,
				"green": 3,
				"blue":  2,
			}},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerOfSet(tt.args.mins); got != tt.want {
				t.Errorf("PowerOfSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 2 example",
			args: args{input: ReadContent(ReadExample())},
			want: 2286,
		},
		{
			name: "Part 2 example",
			args: args{input: ReadContent(ReadInput())},
			want: 67363,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.input); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameLoop(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name     string
		args     args
		gameMins map[string]int
		fail     bool
	}{
		{
			name: "Game 1",
			args: args{game: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			gameMins: map[string]int{
				"red":   4,
				"green": 2,
				"blue":  6,
			},
			fail: false,
		},
		{
			name: "Game 2",
			args: args{game: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
			gameMins: map[string]int{
				"red":   1,
				"green": 3,
				"blue":  4,
			},
			fail: false,
		},
		{
			name: "Game 3",
			args: args{game: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"},
			gameMins: map[string]int{
				"red":   20,
				"green": 13,
				"blue":  6,
			},
			fail: false,
		},
		{
			name: "Game 4",
			args: args{game: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"},
			gameMins: map[string]int{
				"red":   14,
				"green": 3,
				"blue":  15,
			},
			fail: false,
		},
		{
			name: "Game 5",
			args: args{game: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"},
			gameMins: map[string]int{
				"red":   6,
				"green": 3,
				"blue":  2,
			},
			fail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GameLoop(tt.args.game)
			if !reflect.DeepEqual(got, tt.gameMins) {
				t.Errorf("GameLoop() got = %v, want %v", got, tt.gameMins)
			}
			if got1 != tt.fail {
				t.Errorf("GameLoop() got1 = %v, want %v", got1, tt.fail)
			}
		})
	}
}
