package day13

import (
	"encoding/json"
	"os"
	"testing"
)

var p *Puzzle = &Puzzle{}
var exampleInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

func TestParsePackage(t *testing.T) {
	tests := []struct {
		Input    string
		Expected []interface{}
	}{
		{"[1,1,3,1,1]", []interface{}{1, 1, 3, 1, 1}},
		{"[[1],[2,3,4]]", []interface{}{[]int{1}, []int{2, 3, 4}}},
		{"[[[]]]", []interface{}{[]interface{}{[]interface{}{}}}},
		{"[[[],[]]]", []interface{}{[]interface{}{[]interface{}{}, []interface{}{}}}},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", []interface{}{1, []interface{}{2, []interface{}{3, []interface{}{4, []int{5, 6, 7}}}}, 8, 9}},
	}

	for i, tt := range tests {
		got := parsePackage(tt.Input)
		gotByte, _ := json.Marshal(got)
		expectedByte, _ := json.Marshal(tt.Expected)
		gotStr := string(gotByte)
		expectedStr := string(expectedByte)
		if gotStr != expectedStr {
			t.Errorf("[%02d] Expected %q, got %q", i, expectedStr, gotStr)
		}
	}
}

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
	if solution != "5013" {
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
