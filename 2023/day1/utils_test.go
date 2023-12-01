package main

import (
	"testing"
)

func Test_findNumbers(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		startNum int
		endNum   int
		wantErr  bool
	}{
		{
			name:     "2 number line",
			args:     args{line: "1abc2"},
			startNum: 1,
			endNum:   2,
			wantErr:  false,
		},
		{
			name:     "2 numbers middle line",
			args:     args{line: "pqr3stu8vwx"},
			startNum: 3,
			endNum:   8,
			wantErr:  false,
		},
		{
			name:     "Many nymbers",
			args:     args{line: "a1b2c3d4e5f"},
			startNum: 1,
			endNum:   5,
			wantErr:  false,
		},
		{
			name:     "Single number",
			args:     args{line: "pqr7stu"},
			startNum: 7,
			endNum:   7,
			wantErr:  false,
		},
		{
			name:     "Numbers together",
			args:     args{line: "pqr77stu"},
			startNum: 7,
			endNum:   7,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startNum, endNum, err := findNumbers(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("findNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if startNum != tt.startNum {
				t.Errorf("findNumbers() startNum = %v, want %v", startNum, tt.startNum)
			}
			if endNum != tt.endNum {
				t.Errorf("findNumbers() endNum = %v, want %v", endNum, tt.endNum)
			}
		})
	}
}

func Test_concatNumbers(t *testing.T) {
	type args struct {
		startNum int
		endNum   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Add numbers",
			args:    args{startNum: 1, endNum: 5},
			want:    15,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := concatNumbers(tt.args.startNum, tt.args.endNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("concatNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("concatNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumAllNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Add numbers",
			args: args{nums: []int{15, 16}},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAllNumbers(tt.args.nums); got != tt.want {
				t.Errorf("sumAllNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTextToNums(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Replace number text to digit",
			args: args{line: "4nfone5eight"},
			want: "4nfon1e5ei8ght",
		},
		{
			name: "Full numbers string",
			args: args{line: "eightwothree"},
			want: "ei8ghtw2oth3ree",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTextToNums(tt.args.line); got != tt.want {
				t.Errorf("convertTextToNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foo(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 2",
			args: args{file: "input.test2.txt"},
			want: 281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := everything(tt.args.file); got != tt.want {
				t.Errorf("foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
