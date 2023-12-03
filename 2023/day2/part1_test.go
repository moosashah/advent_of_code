package main

import (
	"reflect"
	"testing"
)

func TestIsRoundPossible(t *testing.T) {
	type args struct {
		r      string
		limits map[string]int
	}
	tests := []struct {
		name     string
		args     args
		possible bool
	}{
		{
			name:     "Example input",
			args:     args{r: "3 blue, 4 red", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: true,
		},
		{
			name:     "pass green",
			args:     args{r: "3 blue, 4 red, 13 green", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: true,
		},
		{
			name:     "pass red",
			args:     args{r: "3 blue, 12 red, 13 green", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: true,
		},
		{
			name:     "pass blue",
			args:     args{r: "14 blue, 4 red, 12 green", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: true,
		},
		{
			name:     "fail green",
			args:     args{r: "3 blue, 4 red, 14 green", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: false,
		},
		{
			name:     "fail red",
			args:     args{r: "3 blue, 13 red, 13 green", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: false,
		},
		{
			name:     "fail blue",
			args:     args{r: "15 blue, 4 red, 12 green", limits: map[string]int{"red": 12, "green": 13, "blue": 14}},
			possible: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if possible := IsRoundPossible(tt.args.r, tt.args.limits); possible != tt.possible {
				t.Errorf("IsRoundPossible() = %v, possible %v", possible, tt.possible)
			}
		})
	}
}

func TestGetGameNumber(t *testing.T) {
	type args struct {
		g string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Pull game number",
			args: args{g: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: 1,
		},
		{
			name: "Pull big number",
			args: args{g: "Game 15: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: 15,
		},
		{
			name: "Pull 100",
			args: args{g: "Game 100: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGameNumber(tt.args.g); got != tt.want {
				t.Errorf("GetGameNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playExample(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Example",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playExample(); got != tt.want {
				t.Errorf("playExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitGameAndRounds(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name   string
		args   args
		gameNr int
		rds    []string
	}{
		{
			name:   "Game Possible",
			args:   args{game: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			gameNr: 1,
			rds:    []string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"},
		},
		{
			name:   "Game fail",
			args:   args{game: "Game 2: 30 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			gameNr: 2,
			rds:    []string{"30 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"},
		},
		{
			name:   "Game fail round 2",
			args:   args{game: "Game 1: 3 blue, 4 red; 1 red, 20 green, 6 blue; 2 green"},
			gameNr: 1,
			rds:    []string{"3 blue, 4 red", "1 red, 20 green, 6 blue", "2 green"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SplitGameAndRounds(tt.args.game)
			if got != tt.gameNr {
				t.Errorf("SplitGameAndRounds() got = %v, want %v", got, tt.gameNr)
			}
			if !reflect.DeepEqual(got1, tt.rds) {
				t.Errorf("SplitGameAndRounds() got1 = %v, want %v", got1, tt.rds)
			}
		})
	}
}

func TestIsGamePossible(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name   string
		args   args
		gameNr int
	}{
		{
			name:   "Game Possible",
			args:   args{game: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			gameNr: 1,
		},
		{
			name:   "Game fail",
			args:   args{game: "Game 2: 30 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			gameNr: 0,
		},
		{
			name:   "Game fail round 2",
			args:   args{game: "Game 1: 3 blue, 4 red; 1 red, 20 green, 6 blue; 2 green"},
			gameNr: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGamePossible(tt.args.game); got != tt.gameNr {
				t.Errorf("IsGamePossible() = %v, want %v", got, tt.gameNr)
			}
		})
	}
}

func TestSplitRounds(t *testing.T) {
	type args struct {
		rounds string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Game fail round 2",
			args: args{rounds: "3 blue, 4 red; 1 red, 20 green, 6 blue; 2 green"},
			want: []string{"3 blue, 4 red", "1 red, 20 green, 6 blue", "2 green"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitRounds(tt.args.rounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
