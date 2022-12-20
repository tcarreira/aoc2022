package day20

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `1
2
-3
3
-2
0
4
`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "3"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "5498" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "1623178306"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "3390007892081" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
