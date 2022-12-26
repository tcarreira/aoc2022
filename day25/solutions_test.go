package day25

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

var decimalToSnafuMap = map[int]string{
	1:         "1",
	2:         "2",
	3:         "1=",
	4:         "1-",
	5:         "10",
	6:         "11",
	7:         "12",
	8:         "2=",
	9:         "2-",
	10:        "20",
	15:        "1=0",
	20:        "1-0",
	2022:      "1=11-2",
	12345:     "1-0---0",
	314159265: "1121-1110-1=0",
}

func TestSnafuToDecimal(t *testing.T) {
	for dec, snafu := range decimalToSnafuMap {
		got := snafuToDecimal(snafu)
		if got != dec {
			t.Errorf("Got %d, expected %d", got, dec)
		}
	}
}

func TestDecimalToSnafu(t *testing.T) {
	for dec, snafu := range decimalToSnafuMap {
		got := decimalToSnafu(dec)
		if got != snafu {
			t.Errorf("Got %s, expected %s (dec: %d)", got, snafu, dec)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "2=-1=0"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "2011-=2=-1020-1===-1" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{"", "-"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "-" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
