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

type Node struct {
	visited bool
	isCube  bool
}

func buildMap3D(cubes []Cube) CubesMap {
	myMap := [25][25][25]Node{}

	for _, c := range cubes {
		// if c.x == 0 || c.y == 0 || c.z == 0 {
		// 	fmt.Println("DEBUG:::::", c)
		// }
		myMap[c.x+1][c.y+1][c.z+1].isCube = true // shift all to inside
	}
	return myMap
}

type CubesMap [25][25][25]Node

func (m *CubesMap) dfsSearch(c Cube) {
	m[c.x][c.y][c.z].visited = true
	L := len(m)

	// check neighbors
	for _, p := range []Cube{
		{c.x, c.y, c.z + 1},
		{c.x, c.y, c.z - 1},
		{c.x, c.y + 1, c.z},
		{c.x, c.y - 1, c.z},
		{c.x + 1, c.y, c.z},
		{c.x - 1, c.y, c.z},
	} {
		if p.x < 0 || p.x >= L || p.y < 0 || p.y >= L || p.z < 0 || p.z >= L {
			continue
		}
		if !m[p.x][p.y][p.z].isCube && !m[p.x][p.y][p.z].visited {
			m.dfsSearch(p)
		}
	}
}

func (m *CubesMap) area() int {
	area := 0
	L := len(m)

	for x := range m {
		for y := range m[x] {
			for z := range m[x][y] {
				if !m[x][y][z].isCube {
					continue
				}

				for _, p := range []Cube{
					{x, y, z + 1},
					{x, y, z - 1},
					{x, y + 1, z},
					{x, y - 1, z},
					{x + 1, y, z},
					{x - 1, y, z},
				} {
					if p.x < 0 || p.x >= L || p.y < 0 || p.y >= L || p.z < 0 || p.z >= L {
						continue
					}
					if m[p.x][p.y][p.z].visited {
						area += 1
					}
				}
			}
		}
	}
	return area
}

func (*Puzzle) Part1(input string) string {
	cubes := parseInput(input)
	_, area := buildCubesMap(cubes)
	return fmt.Sprint(area)
}

func (*Puzzle) Part2(input string) string {
	cubes := parseInput(input)
	cubesMap := buildMap3D(cubes)

	cubesMap.dfsSearch(Cube{0, 0, 0})
	area := cubesMap.area()

	return fmt.Sprint(area)
}

func (*Puzzle) Notes() string {
	return "análogo a encontrar saída de labirinto em 3D"
}
