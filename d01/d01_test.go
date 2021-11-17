package d01

import "testing"

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
		{"1", args{"+1, +1, +1"}, 3},
		{"1", args{"+1, +1, -2"}, 0},
		{"1", args{"-1, -2, -3"}, -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickRun1(tt.args.input); got != tt.want {
				t.Errorf("QuickRun1() = %v, want %v", got, tt.want)
			}
		})
	}
}
