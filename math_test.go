package atom

import "testing"

func TestPrecision(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0.001", args{0.001}, 3},
		{"0.002", args{0.002}, 2},
		{"0.01", args{0.01}, 2},
		{"0.1", args{0.1}, 1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Precision(tt.args.a); got != tt.want {
				t.Errorf("Precision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		a float64
		p int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"0.001", args{0.10000003, 2}, 0.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}