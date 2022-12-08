package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

// 30373
// 25512
// 65332
// 33549
// 35390

// Input: 30373
// PD   : OxxOx
// top  : 7

func parseInput(raw string) [][]int {
	grid := [][]int{}

	lines := strings.Split(strings.TrimSpace(raw), "\n")
	for i, line := range lines {
		grid = append(grid, []int{})
		for _, c := range line {
			h, err := strconv.Atoi(string(c))
			if err != nil {
				panic(fmt.Errorf("not parsable: %w", err))
			}
			grid[i] = append(grid[i], h)
		}
	}
	return grid
}

func countVisible(grid [][]bool) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] {
				sum += 1
			}
		}
	}
	return sum
}

func (*Puzzle) Part1(input string) string {
	grid := parseInput(input)
	visibleGrid := checkVisible(grid)
	sum := countVisible(visibleGrid)
	return fmt.Sprint(sum)
}

func checkVisible(exampleGrid [][]int) [][]bool {
	PD := [][]bool{}

	for i, line := range exampleGrid {
		PD = append(PD, []bool{})
		top := -1
		for _, c := range line {
			if c > top {
				top = c
				PD[i] = append(PD[i], true)
			} else {
				PD[i] = append(PD[i], false)
			}
		}
	}

	for i, line := range exampleGrid {
		top := -1
		for j := len(line) - 1; j >= 0; j-- {
			c := line[j]
			if c > top {
				top = c
				PD[i][j] = true
			}
		}
	}

	for i := 0; i <= len(exampleGrid)-1; i++ {
		line := exampleGrid[i]
		top := -1
		for j := 0; j <= len(line)-1; j++ {
			c := exampleGrid[j][i]
			// fmt.Println(i, j, c, top)
			if c > top {
				top = c
				PD[j][i] = true
			}
		}
	}

	for i := len(exampleGrid) - 1; i >= 0; i-- {
		line := exampleGrid[i]
		top := -1
		for j := len(line) - 1; j >= 0; j-- {
			c := exampleGrid[j][i]
			if c > top {
				top = c
				PD[j][i] = true
			}
		}
	}

	return PD
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
