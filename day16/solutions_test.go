package day16

import (
	"fmt"
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "1651"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "2330" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "1707"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "2675" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}

func TestPopLeft(t *testing.T) {
	list := Queue{"a", "b", "c"}
	if fmt.Sprint(list) != "[a b c]" {
		t.Error("not equal 1")
	}

	if list.popLeft() != "a" {
		t.Error("should get A")
	}
	if fmt.Sprint(list) != "[b c]" {
		t.Error("should have less 1")
	}
}

func TestComputeDistanceToAll(t *testing.T) {
	c := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D"},
		"C": {"A"},
		"D": {"B", "E"},
		"E": {"D"},
	}
	nonEmpty := map[string]int{"A": 1, "B": 1, "C": 1, "D": 1, "E": 1}
	got := computeDistanceToAll(c, nonEmpty, "A")
	expected := map[string]int{"B": 1, "C": 1, "D": 2, "E": 3}
	if fmt.Sprint(got) != fmt.Sprint(expected) {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}
func TestComputeValvesDistances(t *testing.T) {
	v := parseInput(exampleInput)
	nonEmpty := findNonEmptyValves(v)
	expected := map[string]map[string]int{
		"AA": {"BB": 1, "CC": 2, "DD": 1, "EE": 2, "HH": 5, "JJ": 2},
		"BB": {"CC": 1, "DD": 2, "EE": 3, "HH": 6, "JJ": 3},
		"CC": {"BB": 1, "DD": 1, "EE": 2, "HH": 5, "JJ": 4},
		"DD": {"BB": 2, "CC": 1, "EE": 1, "HH": 4, "JJ": 3},
		"EE": {"BB": 3, "CC": 2, "DD": 1, "HH": 3, "JJ": 4},
		"HH": {"BB": 6, "CC": 5, "DD": 4, "EE": 3, "JJ": 7},
		"JJ": {"BB": 3, "CC": 4, "DD": 3, "EE": 4, "HH": 7},
	}
	got := computeValvesDistances(v.connections, nonEmpty)
	if fmt.Sprint(got) != fmt.Sprint(expected) {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}

}

func TestPathToState(t *testing.T) {
	s := []string{"AA", "BB", "DD", "CC"}
	got := pathToState(s)
	if got != "BB,CC,DD" {
		t.Error("did not order. got", got)
	}
	if fmt.Sprint(s) != "[AA BB DD CC]" {
		t.Error("modified original :(", fmt.Sprint(s))
	}
}

func TestOverlap(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		Expected bool
	}{
		{"AA", "AA", true},
		{"AA,BB,CC,DD", "EE,FF", false},
		{"EE,FF", "AA,BB,CC,DD", false},
		{"AA,BB,CC", "AA", true},
		{"AA,BB,CC", "BB", true},
		{"AA,BB,CC", "CC", true},
	}

	for i, tt := range tests {
		got := overlap(tt.s1, tt.s2)
		if got != tt.Expected {
			t.Errorf("[%02d] Expected %v, got %v", i, tt.Expected, got)
		}
	}

}
