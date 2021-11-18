package d01

import (
	"testing"
)

func TestQuickRun1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"+1, -2, +3, +1"}, 3},
		{"2", args{"+1, +1, +1"}, 3},
		{"3", args{"+1, +1, -2"}, 0},
		{"4", args{"-1, -2, -3"}, -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickRun1(tt.args.input); got != tt.want {
				t.Errorf("QuickRun1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickRun2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{input: []string{"+1", "-2", "+3", "+1"}}, 2},
		{"2", args{input: []string{"+1", "-1"}}, 0},
		{"3", args{input: []string{"+3", "+3", "+4", "-2", "-4"}}, 10},
		{"4", args{input: []string{"-6", "+3", "+8", "+5", "-6"}}, 5},
		{"5", args{input: []string{"+7", "+7", "-2", "-7", "-4"}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickRun2(tt.args.input); got != tt.want {
				t.Errorf("QuickRun2() = %v, want %v", got, tt.want)
			}
		})
	}
}
