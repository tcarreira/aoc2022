package day18

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

type Cube struct {
	x int
	y int
	z int
}

func parseInput(raw string) []Cube {
	cubes := []Cube{}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		cubes = append(cubes, Cube{x, y, z})
	}
	return cubes
}

func buildCubesMap(cubes []Cube) (map[Cube]bool, int) {
	cubesMap := map[Cube]bool{}
	area := 0
	for _, cube := range cubes {
		cubesMap[cube] = true
		area += 6

		// check neighbors
		if cubesMap[Cube{cube.x, cube.y, cube.z + 1}] { // up
			area -= 2
		}
		if cubesMap[Cube{cube.x, cube.y, cube.z - 1}] { // down
			area -= 2
		}
		if cubesMap[Cube{cube.x + 1, cube.y, cube.z}] { // front
			area -= 2
		}
		if cubesMap[Cube{cube.x - 1, cube.y, cube.z}] { // back
			area -= 2
		}
		if cubesMap[Cube{cube.x, cube.y + 1, cube.z}] { // left
			area -= 2
		}
		if cubesMap[Cube{cube.x, cube.y - 1, cube.z}] { // right
			area -= 2
		}
	}
	return cubesMap, area
}

func (*Puzzle) Part1(input string) string {
	cubes := parseInput(input)
	_, area := buildCubesMap(cubes)
	// buildCubesMap(cubes) -> map[P]bool, area
	return fmt.Sprint(area)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
