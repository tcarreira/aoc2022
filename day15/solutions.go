package day15

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

// S = (3,2)
// B = (6,5)
// dist = 3+3 = 6
// yDist = 7-2 = 5
// xDist = 6-5 = 1
// for 2,3,4
// . . . . . . . . . .
// . . . . . . . . . .
// . . . S . . . . . .
// . . . . . B . . . .
// . . . . . . . . . .
// . . . . . . . . . .
// . . . . . . . . . .
// . . # # # . . . . .  <==== 7
// . . . . . . . . . .
// . . . . . . . . . .

type P struct {
	x int
	y int
}
type Sensor struct {
	position P
	beacon   P
}

func parseInput(raw string) []Sensor {
	sensors := []Sensor{}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		var Sx, Sy, Bx, By int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &Sx, &Sy, &Bx, &By)
		sensors = append(sensors, Sensor{P{Sx, Sy}, P{Bx, By}})
	}
	return sensors
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(a, b P) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func findCoverage(sensors []Sensor, y int) int {
	coverage := make(map[int]bool)
	for _, sensor := range sensors {
		dist := manhattan(sensor.position, sensor.beacon)
		yDist := abs(y - sensor.position.y)
		xDist := dist - yDist // -2
		for x := sensor.position.x - xDist; x <= sensor.position.x+xDist; x++ {
			if sensor.beacon.y != y || x != sensor.beacon.x {
				coverage[x] = true
			}
		}
	}

	return len(coverage)
}

func part1(input string, y int) string {
	// parseInput() -> [](Sensor,Beacon)
	// findCoverage(y=2000000)
	//   - for pair(S,B) - dist()
	//     - dist(S,x) <= dist
	sensors := parseInput(input)
	solution := findCoverage(sensors, y)
	return fmt.Sprint(solution)
}

func (*Puzzle) Part1(input string) string {
	return part1(input, 2000000)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
