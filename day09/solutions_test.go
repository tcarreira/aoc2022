package day09

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "13"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "6391" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	example2 := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "1"},
		{example2, "36"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "2593" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
