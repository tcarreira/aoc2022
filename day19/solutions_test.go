package day19

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian. 
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "33"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "1616" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := p.Part1(exampleInput)
		expected := "0"
		if got != expected {
			b.Errorf("Expected %q, got %q", expected, got)
		}
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
