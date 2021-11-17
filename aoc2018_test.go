package main

import "testing"

func Test_greeting(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"dummy", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greeting(); got != tt.want {
				t.Errorf("greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}
