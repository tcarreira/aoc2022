package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

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

func checkVisible(grid [][]int) [][]bool {
	visibleTrees := [][]bool{}

	// left - right (build)
	for i, line := range grid {
		visibleTrees = append(visibleTrees, []bool{})
		top := -1
		for _, c := range line {
			if c > top {
				top = c
				visibleTrees[i] = append(visibleTrees[i], true)
			} else {
				visibleTrees[i] = append(visibleTrees[i], false)
			}
		}
	}

	// right - left
	for i, line := range grid {
		top := -1
		for j := len(line) - 1; j >= 0; j-- {
			c := line[j]
			if c > top {
				top = c
				visibleTrees[i][j] = true
			}
		}
	}

	// down - up
	for i := 0; i <= len(grid)-1; i++ {
		line := grid[i]
		top := -1
		for j := 0; j <= len(line)-1; j++ {
			c := grid[j][i]
			// fmt.Println(i, j, c, top)
			if c > top {
				top = c
				visibleTrees[j][i] = true
			}
		}
	}

	// up - down
	for i := len(grid) - 1; i >= 0; i-- {
		line := grid[i]
		top := -1
		for j := len(line) - 1; j >= 0; j-- {
			c := grid[j][i]
			if c > top {
				top = c
				visibleTrees[j][i] = true
			}
		}
	}

	return visibleTrees
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

func bestScenicView(grid [][]int) int {
	maxScenic := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tree := grid[i][j]
			scenic := 1

			// ->
			count := 0
			for y := j + 1; y < len(grid); y++ {
				count++
				if grid[i][y] >= tree {
					break
				}
			}
			scenic *= count

			// <-
			count = 0
			for y := j - 1; y >= 0; y-- {
				count++
				if grid[i][y] >= tree {
					break
				}
			}
			scenic *= count

			// ^
			count = 0
			for x := i - 1; x >= 0; x-- {
				count++
				if grid[x][j] >= tree {
					break
				}
			}
			scenic *= count

			// v
			count = 0
			for x := i + 1; x < len(grid); x++ {
				count++
				if grid[x][j] >= tree {
					break
				}
			}
			scenic *= count

			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
	}
	return maxScenic
}

func (*Puzzle) Part1(input string) string {
	grid := parseInput(input)
	visibleGrid := checkVisible(grid)
	sum := countVisible(visibleGrid)
	return fmt.Sprint(sum)
}

func (*Puzzle) Part2(input string) string {
	grid := parseInput(input)
	sum := bestScenicView(grid)
	return fmt.Sprint(sum)

}

func (*Puzzle) Notes() string {
	return "programação dinâmica"
}
