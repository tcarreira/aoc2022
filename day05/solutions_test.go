package day05

import (
	"io/ioutil"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPop(t *testing.T) {
	s := Stack{"A", "B", "C"}
	o := s.Pop()
	if o != "C" {
		t.Error("Not the right element")
	}
	if len(s) != 2 {
		t.Error("not right len")
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "CMZ"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := ioutil.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "WSFTMRHPP" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "MCD"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := ioutil.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "GSLCMFBRP" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
