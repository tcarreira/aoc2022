package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

type Pair struct {
	range1 []int
	range2 []int
}

func parseInput(raw string) []Pair {
	pairs := strings.Split(strings.TrimSpace(raw), "\n")

	result := []Pair{}
	for _, pair := range pairs {
		subPairs := strings.Split(pair, ",")
		subPair1 := strings.Split(subPairs[0], "-")
		subPair2 := strings.Split(subPairs[1], "-")

		subPair1Int1, _ := strconv.Atoi(subPair1[0])
		subPair1Int2, _ := strconv.Atoi(subPair1[1])
		subPair2Int1, _ := strconv.Atoi(subPair2[0])
		subPair2Int2, _ := strconv.Atoi(subPair2[1])

		result = append(result, Pair{
			range1: []int{subPair1Int1, subPair1Int2},
			range2: []int{subPair2Int1, subPair2Int2},
		})
	}
	return result
}

func isFullyContained(range1, range2 []int) int {
	if range1[0] <= range2[0] && range1[1] >= range2[1] {
		return 1
	}
	if range1[0] >= range2[0] && range1[1] <= range2[1] {
		return 1
	}
	return 0
}

func (*Puzzle) Part1(input string) string {
	// lista de pares
	// sub-par: intervalo inclusivo
	// em quantos pares existe sobreposição completa
	parsedInput := parseInput(input)

	count := 0
	for _, pair := range parsedInput {
		count += isFullyContained(pair.range1, pair.range2)
	}

	return fmt.Sprint(count)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
