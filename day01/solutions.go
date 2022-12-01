package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct{}

type ElfsLists [][]int

func parseInput(raw string) ElfsLists {
	elfLists := strings.Split(raw, "\n\n")

	result := ElfsLists{}
	for _, elfList := range elfLists {
		itemsStr := strings.Split(elfList, "\n")

		elfItems := []int{}
		for _, itemStr := range itemsStr {
			itemInt, _ := strconv.Atoi(itemStr)
			elfItems = append(elfItems, itemInt)
		}
		result = append(result, elfItems)
	}

	return result
}

func (*Puzzle) Part1(input string) string {
	parsedInput := parseInput(input)

	maxSum := 0
	for _, elfList := range parsedInput {
		elfCalSum := 0
		for _, item := range elfList {
			elfCalSum += item
		}

		if elfCalSum > maxSum {
			maxSum = elfCalSum
		}
	}

	return fmt.Sprint(maxSum)
}

func (*Puzzle) Part2(input string) string {
	parsedInput := parseInput(input)

	var elfCalSums []int
	for _, elfList := range parsedInput {
		elfCalSum := 0
		for _, item := range elfList {
			elfCalSum += item
		}

		elfCalSums = append(elfCalSums, elfCalSum)
	}

	sort.Ints(elfCalSums)

	L := len(elfCalSums)
	top3Sum := elfCalSums[L-1] + elfCalSums[L-2] + elfCalSums[L-3]
	return fmt.Sprint(top3Sum)
}

func (*Puzzle) Notes() string {
	return "primeiro problema"
}
