package day05

import (
	"strconv"
	"strings"
)

type Puzzle struct{}

type Stack []string
type Stacks []Stack
type Move struct {
	craneCount int
	fromStack  int
	toStack    int
}

type ParsedInput struct {
	stacks *Stacks
	moves  []Move
}

func parseStacks(stacksStr string) *Stacks {
	lines := strings.Split(stacksStr, "\n")
	width := len(lines[0])

	stacks := make(Stacks, width/4+1)
	for l := len(lines) - 2; l >= 0; l-- {
		column := -1
		for c := 1; c < width; c += 4 {
			column++
			if string(lines[l][c]) == " " {
				continue
			}
			stacks[column] = append(stacks[column], string(lines[l][c]))
		}
	}
	return &stacks
}

func parseMoves(movesStr string) []Move {
	lines := strings.Split(movesStr, "\n")
	moves := []Move{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		craneCount, _ := strconv.Atoi(parts[1])
		fromStack, _ := strconv.Atoi(parts[3])
		toStack, _ := strconv.Atoi(parts[5])

		moves = append(moves, Move{
			craneCount: craneCount,
			fromStack:  fromStack,
			toStack:    toStack,
		})
	}
	return moves
}

func parseInput(raw string) ParsedInput {
	parts := strings.Split(strings.TrimRight(raw, "\n"), "\n\n")

	return ParsedInput{
		stacks: parseStacks(parts[0]),
		moves:  parseMoves(parts[1]),
	}
}

func (s *Stack) Pop() string {
	out := (*s)[len(*s)-1]
	(*s)[len(*s)-1] = ""
	*s = (*s)[:len(*s)-1]
	return out
}

func (s *Stack) Push(c string) {
	(*s) = append((*s), c)
}

func (s *Stacks) applyMoves(move Move) {
	for i := 0; i < move.craneCount; i++ {
		crane := (*s)[move.fromStack-1].Pop()
		(*s)[move.toStack-1].Push(crane)
	}
}

func (s *Stacks) applyMoves2(move Move) {
	movingCrates := (*s)[move.fromStack-1][len((*s)[move.fromStack-1])-move.craneCount:]
	(*s)[move.fromStack-1] = (*s)[move.fromStack-1][:len((*s)[move.fromStack-1])-move.craneCount]
	(*s)[move.toStack-1] = append((*s)[move.toStack-1], movingCrates...)
}

func (*Puzzle) Part1(input string) string {
	parsedInput := parseInput(input)

	for _, move := range parsedInput.moves {
		parsedInput.stacks.applyMoves(move)
	}

	result := ""
	for _, stack := range *parsedInput.stacks {
		result += stack.Pop()
	}
	return result
}

func (*Puzzle) Part2(input string) string {
	parsedInput := parseInput(input)

	for _, move := range parsedInput.moves {
		parsedInput.stacks.applyMoves2(move)
	}

	result := ""
	for _, stack := range *parsedInput.stacks {
		result += stack.Pop()
	}
	return result
}

func (*Puzzle) Notes() string {
	return ""
}
