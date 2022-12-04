package day04

import (
	"io/ioutil"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "2"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := ioutil.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "485" {
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

	data, _ := ioutil.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "-" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
