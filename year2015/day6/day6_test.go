package day6

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"", "1000000", "turn on 0,0 through 999,999"},
		{"", "1000000", "turn on 0,0 through 999,999\n\n"},
		{"", "1000", "toggle 0,0 through 999,0"},
		{"", "0", "turn off 499,499 through 500,500"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"", "1", "turn on 0,0 through 0,0"},
		{"", "1000000", "turn on 0,0 through 999,999"},
		{"", "1000000", "turn on 0,0 through 999,999\n\n"},
		{"", "2000000", "toggle 0,0 through 999,999"},
		{"", "2000", "toggle 0,0 through 999,0"},
		{"", "0", "turn off 499,499 through 500,500"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Part2(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestCountLit(t *testing.T) {
	lights := make([][]bool, 3)
	for i := range lights {
		lights[i] = make([]bool, 3)
		for j := 0; j < 3; j++ {
			if j == i {
				lights[i][j] = true
			}
		}
	}
	var tests = []struct {
		name     string
		expected int
		given    [][]bool
	}{
		{"", 3, lights},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := countLit(tt.given)
			if actual != tt.expected {
				t.Errorf("countLit(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}
