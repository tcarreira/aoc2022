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

func isAllDifferent(w string) bool {
	// 0 1 2 3 4 5 6 7 8 9
	// 0 - - - - - - - - -
	//   1 - - - - - - - -
	//     2 - - - - - - -
	//       3 - - - - - -
	//         4 - - - - -
	// ...
	//                 8

	for i, c1 := range w[:len(w)-1] {
		for _, c2 := range w[i+1:] {
			if c1 == c2 {
				return false
			}
		}
	}
	return true
}

func (*Puzzle) Part2(input string) string {
	input = strings.TrimSpace(input)
	for i := 0; i < len(input)-14; i++ {
		w := input[i : i+14]

		if isAllDifferent(w) {
			return fmt.Sprint(i + 14)
		}
	}
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
