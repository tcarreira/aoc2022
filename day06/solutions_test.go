package day06

import (
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "7"},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "5"},
		{"nppdvjthqldpwncqszvftbrmjlhg", "6"},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "10"},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "11"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "1855" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "19"},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "23"},
		{"nppdvjthqldpwncqszvftbrmjlhg", "23"},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "29"},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "26"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "3256" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
