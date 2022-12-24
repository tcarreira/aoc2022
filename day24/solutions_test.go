package day24

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var smallerExample = `
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#
`
var exampleInput = `
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
`

func TestWrap(t *testing.T) {
	tests := []struct {
		n        int
		nmin     int
		nmax     int
		Expected int
	}{
		// 1 <= n <= 10
		{1, 1, 10, 1},
		{2, 1, 10, 2},
		{10, 1, 10, 10},
		{11, 1, 10, 1},
		{19, 1, 10, 9},
		{20, 1, 10, 10},
		{21, 1, 10, 1},
		{0, 1, 10, 10},
		{-1, 1, 10, 9},
		{-9, 1, 10, 1},
		{-10, 1, 10, 10},
		{-19, 1, 10, 1},
		{-20, 1, 10, 10},
		{-21, 1, 10, 9},
		// 1 <= n <= 5
		{1, 1, 5, 1},
		{2, 1, 5, 2},
		{5, 1, 5, 5},
		{1, 1, 5, 1},
		{0, 1, 5, 5},
	}

	for i, tt := range tests {
		got := wrap(tt.n, tt.nmin, tt.nmax)
		if got != tt.Expected {
			t.Errorf("[%02d] Expected %d, got %d", i, tt.Expected, got)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{smallerExample, "10"},
		{exampleInput, "18"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "281" {
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
