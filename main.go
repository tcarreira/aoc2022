package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/tcarreira/aoc2022/day01"
	"github.com/tcarreira/aoc2022/day02"
	"github.com/tcarreira/aoc2022/day03"
	"github.com/tcarreira/aoc2022/day04"
	"github.com/tcarreira/aoc2022/day05"
	"github.com/tcarreira/aoc2022/day06"
	// ci:importDay
)

const nTries int = 10

type DayAOC interface {
	Part1(input string) string
	Part2(input string) string
	Notes() string
}

func main() {
	fmt.Println("######################################################")
	fmt.Println("Solving problems for the https://adventofcode.com/2022")
	fmt.Println("######################################################")

	fmt.Println()
	fmt.Println(" Day | Part1      ( time ms ) | Part2      ( time ms ) |")
	fmt.Println("-----+------------------------+------------------------+")

	aocDays := []DayAOC{
		// existing days:
		&day01.Puzzle{},
		&day02.Puzzle{},
		&day03.Puzzle{},
		&day04.Puzzle{},
		&day05.Puzzle{},
		&day06.Puzzle{},
	} // ci:addNewDay
	for day0, aocDay := range aocDays {
		day := day0 + 1
		inputBytes, err := ioutil.ReadFile(fmt.Sprintf("day%02d/input.txt", day))
		if err != nil {
			fmt.Println("Error reading input file", err)
			continue
		}
		input := string(inputBytes)

		// Benchmarking part1
		startTime := time.Now()
		part1Solution := ""
		for i := 0; i < nTries; i++ {
			part1Solution = aocDay.Part1(input)
		}
		part1Time := time.Now().Sub(startTime)

		// Benchmarking part2
		startTime = time.Now()
		part2Solution := ""
		for i := 0; i < nTries; i++ {
			part2Solution = aocDay.Part2(input)
		}
		part2Time := time.Now().Sub(startTime)

		// Printing results
		fmt.Printf("  %02d | %-10s (%9.1f) | %-10s (%9.1f) | %s\n", day,
			part1Solution, float64(part1Time/time.Microsecond)/float64(nTries)/1000,
			part2Solution, float64(part2Time/time.Microsecond)/float64(nTries)/1000,
			aocDay.Notes(),
		)
	}
}
