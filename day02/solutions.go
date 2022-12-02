package day02

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

type Round struct {
	playerA string
	playerB string
}
type ParsedInput []Round

func parseInput(raw string) ParsedInput {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	result := ParsedInput{}
	for _, line := range lines {
		r := Round{
			playerA: string(line[0]),
			playerB: string(line[2]),
		}
		result = append(result, r)
	}

	return result
}

func getRoundPoints(r Round) int {
	// result for playerB
	if r.playerA == "A" {
		if r.playerB == "X" {
			return 1 + 3
		} else if r.playerB == "Y" {
			return 2 + 6
		} else if r.playerB == "Z" {
			return 3 + 0
		}
	} else if r.playerA == "B" {
		if r.playerB == "X" {
			return 1 + 0
		} else if r.playerB == "Y" {
			return 2 + 3
		} else if r.playerB == "Z" {
			return 3 + 6
		}
	} else if r.playerA == "C" {
		switch r.playerB {
		case "X":
			return 1 + 6
		case "Y":
			return 2 + 0
		case "Z":
			return 3 + 3
		}
	}

	panic("Unreachable code")
}

func (*Puzzle) Part1(input string) string {
	// 1. parseInput
	// 4. percorrer rodadas
	//    4.1. fazer o cálculo  --> getRoundPoints(plA string, plB string) int
	//         2. conversão B->Pedra...
	//         3. conversão X->Pedra...
	parsedInput := parseInput(input)

	gamePoints := 0
	for _, r := range parsedInput {
		gamePoints += getRoundPoints(r)
	}

	return fmt.Sprint(gamePoints)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
