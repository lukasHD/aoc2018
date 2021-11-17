package commons

import (
	"reflect"
	"testing"
)

func TestLoadLines(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"oneLine", args{"test_input/oneLine.txt"}, []string{"first Line!"}},
		{"TwoLine", args{"test_input/twoLine.txt"}, []string{"first Line!", "second Line!"}},
		{"TwoLineEmpty", args{"test_input/twoLineEmpty.txt"}, []string{"first Line!"}},
		{"ThreeLine", args{"test_input/threeLine.txt"}, []string{"first", "", "third"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadLines(tt.args.fname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
