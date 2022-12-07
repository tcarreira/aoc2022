package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	inputFileFormat = "day%02d/input.txt"
	cacheFileFormat = "day%02d/.cache.json"
)

func calculatePuzzleStats(day int, aocDay DayAOC) (PuzzleStats, error) {
	pSstats := PuzzleStats{}

	inputBytes, err := os.ReadFile(fmt.Sprintf(inputFileFormat, day))
	if err != nil {
		return pSstats, fmt.Errorf("could not read input file %w", err)
	}
	input := string(inputBytes)

	// Benchmarking part1
	startTime := time.Now()
	for i := 0; i < Repeats; i++ {
		pSstats.Results.Part1 = aocDay.Part1(input)
	}
	pSstats.Timing.Part1 = time.Since(startTime) / time.Duration(Repeats)

	// Benchmarking part2
	startTime = time.Now()
	for i := 0; i < Repeats; i++ {
		pSstats.Results.Part2 = aocDay.Part2(input)
	}
	pSstats.Timing.Part2 = time.Since(startTime) / time.Duration(Repeats)
	return pSstats, nil
}

func parseCachedResults(day int, aocDay DayAOC) (PuzzleStats, error) {
	pSstats := PuzzleStats{}

	data, err := os.ReadFile(fmt.Sprintf(cacheFileFormat, day))
	if err != nil {
		return pSstats, err
	}
	err = json.Unmarshal(data, &pSstats)
	return pSstats, err
}

func cachePuzzleStats(p PuzzleStats, day int) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf(cacheFileFormat, day), data, 0644)
	return err
}

func buildPuzzleStats(day int, aocDay DayAOC) (PuzzleStats, error) {
	if UseCachedResult {
		pStats, err := parseCachedResults(day, aocDay)
		if !os.IsNotExist(err) {
			return pStats, nil
		}

		pStats, err = calculatePuzzleStats(day, aocDay)
		if err == nil && pStats.Results.Part1 != "-" && pStats.Results.Part2 != "-" {
			err = cachePuzzleStats(pStats, day)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not write cache: %v\n", err)
				err = nil
			}
		}
		return pStats, err
	}
	return calculatePuzzleStats(day, aocDay)
}
