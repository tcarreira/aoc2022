package day12

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "31"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "520" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "29"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "508" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
