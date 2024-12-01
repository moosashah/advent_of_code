package main

import (
	"testing"
)

func Test_calculateStrength(t *testing.T) {
	type args struct {
		cards string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one pair",
			args: args{cards: "32T3K"},
			want: 1,
		},
		{
			name: "two pair",
			args: args{cards: "KK677"},
			want: 2,
		},
		{
			name: "three pair",
			args: args{cards: "KKK47"},
			want: 3,
		},
		{
			name: "four pair",
			args: args{cards: "77677"},
			want: 5,
		},
		{
			name: "five pair",
			args: args{cards: "KKKKK"},
			want: 6,
		},
		{
			name: "full",
			args: args{cards: "KK777"},
			want: 4,
		},
		{
			name: "all uniques",
			args: args{cards: "KQT87"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateStrength(tt.args.cards); got != tt.want {
				t.Errorf("calculateStrength() = %v, want %v", got, tt.want)
			}
		})
	}
}
