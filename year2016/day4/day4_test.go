package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{input: `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`}, "1514"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.input); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	type args struct {
		input string
		n     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"qzmtzixmtkozyivhz", 343}, "veryencryptedname"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rotate(tt.args.input, tt.args.n); got != tt.want {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
