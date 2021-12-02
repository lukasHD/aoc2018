package d02

import (
	"testing"
)

func Test_getRepeatCountLetters(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{"1", args{"abcdef"}, false, false},
		{"2", args{"bababc"}, true, true},
		{"3", args{"abbcde"}, true, false},
		{"4", args{"abcccd"}, false, true},
		{"5", args{"aabcdd"}, true, false},
		{"6", args{"abcdee"}, true, false},
		{"7", args{"ababab"}, false, true},
		{"8", args{""}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getRepeatCountLetters(tt.args.row)
			if got != tt.want {
				t.Errorf("getRepeatCountLetters() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRepeatCountLetters() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_commonCharacters(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{"1", args{"abcde", "axcye"}, 2, "ace"},
		{"1", args{"fghij", "fguij"}, 1, "fgij"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := commonCharacters(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("commonCharacters() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("commonCharacters() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
