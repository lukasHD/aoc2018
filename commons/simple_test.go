package commons

import (
	"testing"
)

//Test for correct sum in sum
func TestSum1(t *testing.T) {
	a := 11
	b := 22
	sum, err := MySum(a, b)
	if sum != 33 || err != nil {
		t.Fatalf("Not Match")
	}
}

func TestSum2(t *testing.T) {
	a := 10
	b := 22
	expect := a + b
	sum, err := MySum(a, b)
	if sum != expect || err != nil {
		t.Fatalf("No Match! expect: %d got: %d", expect, sum)
	}
}

func TestMySum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "bla", args: args{1, 2}, want: 3, wantErr: false},
		{"blu", args{2, 2}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MySum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
