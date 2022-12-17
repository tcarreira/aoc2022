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

func minDuration(times []time.Duration) time.Duration {
	var minDur time.Duration = times[0]
	for _, d := range times {
		if d < minDur {
			minDur = d
		}
	}
	return minDur
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

func buildCachedPuzzleStats(day int, aocDay DayAOC) (PuzzleStats, error) {
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
