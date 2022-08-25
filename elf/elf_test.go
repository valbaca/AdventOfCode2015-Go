package elf

import "testing"

func TestMod(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{3, 10}, 3},
		{"", args{-3, 10}, 7},
		{"", args{-3, -10}, -3},
		{"", args{3, -10}, -7},
		{"", args{0, -10}, 0},
		{"", args{0, 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mod(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}
