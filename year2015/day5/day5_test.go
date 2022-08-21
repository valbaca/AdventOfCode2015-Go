package day5

import "testing"

func TestPart1(t *testing.T) {

}

func TestIsNice(t *testing.T) {

	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		{"ex1", true, "ugknbfddgicrmopn"},
		{"ex2", true, "aaa"},
		{"ex3", false, "jchzalrnumimnmhp"},
		{"ex4", false, "haegwjzuvuyypxyu"},
		{"ex5", false, "dvszwmarrgswjxmb"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := IsNice(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestIsVowel(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    rune
	}{
		{"", true, 'a'},
		{"", true, 'e'},
		{"", true, 'i'},
		{"", true, 'o'},
		{"", true, 'u'},
		{"", false, 'b'},
		{"", false, 'c'},
		{"", false, 'd'},
		{"", false, 'f'},
		{"", false, 'g'},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := IsVowel(tt.given)
			if actual != tt.expected {
				t.Errorf("IsVowel(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestIsNaughtyPair(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		givenP   rune
		givenR   rune
	}{
		{"", true, 'a', 'b'},
		{"", true, 'c', 'd'},
		{"", true, 'p', 'q'},
		{"", true, 'x', 'y'},
		{"", false, 'x', 's'},
		{"", false, 'b', 'a'},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := IsNaughtyPair(tt.givenP, tt.givenR)
			if actual != tt.expected {
				t.Errorf("(%q,%q): expected %v, actual %v", tt.givenP, tt.givenR, tt.expected, actual)
			}

		})
	}

}

func TestIsNicer(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		{"", true, "qjhvhtzxzqqjkmpb"},
		{"", true, "xxyxx"},
		{"", false, "uurcxstgmygtbstg"},
		{"", false, "ieodomkazucvgmuy"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := IsNicer(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestHasPair(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		{"", false, "aaa"},
		{"", true, "xyxy"},
		{"", true, "aabcdefgaa"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := HasPair(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}

func TestHasDouble(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		{"", true, "xyx"},
		{"", true, "abcdefeghi"},
		{"", true, "aaa"},
		{"", false, "aa"},
		{"", false, "xyz"},
		{"", false, "xyyx"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := HasDouble(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}
