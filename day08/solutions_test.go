package day08

import (
	"encoding/json"
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `30373
25512
65332
33549
35390
`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "21"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "1779" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "8"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "172224" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}

func TestPart3(t *testing.T) {

	var exampleGrid = [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	var resultGrid = [][]bool{
		{true, true, true, true, true},
		{true, true, true, false, true},
		{true, true, false, true, true},
		{true, false, true, false, true},
		{true, true, true, true, true},
	}

	solution := checkVisible(exampleGrid)

	expected, _ := json.Marshal(resultGrid)
	got, _ := json.Marshal(solution)

	if string(got) != string(expected) {
		t.Errorf("Expected %q, got %q", string(expected), string(got))
	}
}
