package day03

import (
	"io/ioutil"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestFindRepeatedItem(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		Expected string
	}{
		{"abcdz", "efghz", "z"},
		{"abcdz", "apoiy", "a"},
		{"abcdzy", "lpoDry", "y"},
	}

	for i, tt := range tests {
		solution := findRepeatedItem(tt.s1, tt.s2)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}
}

func TestConvertItemToPriority(t *testing.T) {
	tests := []struct {
		c        string
		Expected int
	}{
		{"a", 1},
		{"z", 26},
		{"A", 27},
		{"Z", 52},
	}

	for i, tt := range tests {
		solution := convertItemToPriority(tt.c)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %d, got %d", i, tt.Expected, solution)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "157"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := ioutil.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "7727" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "70"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := ioutil.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "2609" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
