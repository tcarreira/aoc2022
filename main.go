package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"syscall"
	"time"

	"github.com/tcarreira/aoc2022/day01"
	"github.com/tcarreira/aoc2022/day02"
	"github.com/tcarreira/aoc2022/day03"
	"github.com/tcarreira/aoc2022/day04"
	"github.com/tcarreira/aoc2022/day05"
	"github.com/tcarreira/aoc2022/day06"
	"github.com/tcarreira/aoc2022/day07"
	"github.com/tcarreira/aoc2022/day08"
	"github.com/tcarreira/aoc2022/day09"
	"github.com/tcarreira/aoc2022/day10"
	"github.com/tcarreira/aoc2022/day11"
	"github.com/tcarreira/aoc2022/day12"
	"github.com/tcarreira/aoc2022/day13"
	"github.com/tcarreira/aoc2022/day14"
	"github.com/tcarreira/aoc2022/day15"
	"github.com/tcarreira/aoc2022/day16"
	"github.com/tcarreira/aoc2022/day17"
	"github.com/tcarreira/aoc2022/day18"
	"github.com/tcarreira/aoc2022/day19"
	"github.com/tcarreira/aoc2022/day20"
	"github.com/tcarreira/aoc2022/day21"
	"github.com/tcarreira/aoc2022/day22"
	"github.com/tcarreira/aoc2022/day23"
	"github.com/tcarreira/aoc2022/day24"
	"github.com/tcarreira/aoc2022/day25"
	// ci:importDay
)

var (
	Repeats         int  = 3
	UseCachedResult bool = false
	RuntimeProfile  bool = false
)

func init() {
	if repeatsStr := os.Getenv("AOC_REPEATS"); repeatsStr != "" {
		repeats, err := strconv.Atoi(repeatsStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "AOC_REPEATS could not be parsed into int: %v\n", err)
		} else {
			Repeats = repeats
		}
	}

	if useCachedStr := os.Getenv("AOC_USE_CACHED_RESULT"); useCachedStr != "" {
		useCached, err := strconv.ParseBool(useCachedStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "AOC_USE_CACHED_RESULT could not be parsed into bool: %v\n", err)
		} else {
			UseCachedResult = useCached
		}
	}

	if doProfileStr := os.Getenv("AOC_PROFILE"); doProfileStr != "" {
		doProfile, err := strconv.ParseBool(doProfileStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "AOC_PROFILE could not be parsed into bool: %v\n", err)
		} else {
			RuntimeProfile = doProfile
		}
	}
}

func handleSignals(retCode int, handle func()) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-sigc
		handle()
		os.Exit(retCode)
	}()
}

func setupProfiling() func() {
	if !RuntimeProfile {
		return func() {}
	}

	handleSignals(1, func() {
		fmt.Println("Stopping CPU profile")
		pprof.StopCPUProfile()
	})

	f, err := os.Create("aoc2022.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	return pprof.StopCPUProfile
}

type DayAOC interface {
	Part1(input string) string
	Part2(input string) string
	Notes() string
}

type PuzzleStats struct {
	Results struct {
		Part1 string
		Part2 string
	}
	Timing struct {
		Part1 time.Duration
		Part2 time.Duration
	}
}

func calculatePuzzleStats(day int, aocDay DayAOC) (PuzzleStats, error) {
	pSstats := PuzzleStats{}

	inputBytes, err := os.ReadFile(fmt.Sprintf(inputFileFormat, day))
	if err != nil {
		return pSstats, fmt.Errorf("could not read input file %w", err)
	}
	input := string(inputBytes)

	// Benchmarking part1
	times := []time.Duration{}
	for i := 0; i < Repeats; i++ {
		startTime := time.Now()
		pSstats.Results.Part1 = aocDay.Part1(input)
		times = append(times, time.Since(startTime))
	}
	pSstats.Timing.Part1 = minDuration(times)

	// Benchmarking part2
	times = []time.Duration{}
	for i := 0; i < Repeats; i++ {
		startTime := time.Now()
		pSstats.Results.Part2 = aocDay.Part2(input)
		times = append(times, time.Since(startTime))
	}
	pSstats.Timing.Part2 = minDuration(times)
	return pSstats, nil
}

func main() {
	d := setupProfiling()
	defer d()

	fmt.Println("######################################################")
	fmt.Println("Resolvendo os puzzles do https://adventofcode.com/2022")
	fmt.Println("######################################################")

	fmt.Println()
	fmt.Println(" Dia | Parte 1        (time ms) | Parte 2        (time ms) |")
	fmt.Println("-----+--------------------------+----------------------------+")

	aocDays := []DayAOC{
		// existing days:
		&day01.Puzzle{},
		&day02.Puzzle{},
		&day03.Puzzle{},
		&day04.Puzzle{},
		&day05.Puzzle{},
		&day06.Puzzle{},
		&day07.Puzzle{},
		&day08.Puzzle{},
		&day09.Puzzle{},
		&day10.Puzzle{},
		&day11.Puzzle{},
		&day12.Puzzle{},
		&day13.Puzzle{},
		&day14.Puzzle{},
		&day15.Puzzle{},
		&day16.Puzzle{},
		&day17.Puzzle{},
		&day18.Puzzle{},
		&day19.Puzzle{},
		&day20.Puzzle{},
		&day21.Puzzle{},
		&day22.Puzzle{},
		&day23.Puzzle{},
		&day24.Puzzle{},
		&day25.Puzzle{},
	} // ci:addNewDay
	for day0, aocDay := range aocDays {
		day := day0 + 1

		var puzzleStats PuzzleStats
		var err error
		if UseCachedResult {
			puzzleStats, err = buildCachedPuzzleStats(day, aocDay)
		} else {
			puzzleStats, err = calculatePuzzleStats(day, aocDay)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error building PuzzleStats for day %d (cached=%v): %v\n", day, UseCachedResult, err)
			continue
		}

		// Printing results
		if day == 25 { // day 25 has only one part
			fmt.Printf("  %02d | %-41s (%7.1f) | %s\n", day,
				puzzleStats.Results.Part1, float64(puzzleStats.Timing.Part1/time.Microsecond)/1000,
				aocDay.Notes(),
			)
		} else {
			fmt.Printf("  %02d | %-14s (%7.1f) | %-14s (%7.1f) | %s\n", day,
				puzzleStats.Results.Part1, float64(puzzleStats.Timing.Part1/time.Microsecond)/1000,
				puzzleStats.Results.Part2, float64(puzzleStats.Timing.Part2/time.Microsecond)/1000,
				aocDay.Notes(),
			)
		}
	}

	// Special prints
	// still print aoc day 10 drawing when using cached result
	if _, ok := os.LookupEnv("AOC_DAY10"); ok && UseCachedResult {
		fmt.Println()
		calculatePuzzleStats(10, &day10.Puzzle{})
	}
}
