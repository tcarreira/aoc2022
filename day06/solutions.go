package day06

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

func (*Puzzle) Part1(input string) string {
	input = strings.TrimSpace(input)
	for i := 0; i < len(input)-4; i++ {
		w := input[i : i+4]

		if w[0] != w[1] && w[0] != w[2] && w[0] != w[3] && w[1] != w[2] && w[1] != w[3] && w[2] != w[3] {
			return fmt.Sprint(i + 4)
		}
	}
	return "-"
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
