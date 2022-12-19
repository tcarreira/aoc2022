package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

type CaveMap [][]bool

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (m CaveMap) draw(p1, p2 []string) {
	xP1, _ := strconv.Atoi(p1[0])
	yP1, _ := strconv.Atoi(p1[1])
	xP2, _ := strconv.Atoi(p2[0])
	yP2, _ := strconv.Atoi(p2[1])

	if xP1 == xP2 {
		// draw vertical
		for i := min(yP1, yP2); i <= max(yP1, yP2); i++ {
			m[i][xP1] = true
		}
	} else if yP1 == yP2 {
		// draw horizontal
		for i := min(xP1, xP2); i <= max(xP1, xP2); i++ {
			m[yP1][i] = true
		}
	} else {
		panic("not vertica/horizontal line")
	}
}

func parseInput(raw string) CaveMap {
	myMap := make(CaveMap, 1000)
	for i := range myMap {
		myMap[i] = make([]bool, 1000)
	}

	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		points := strings.Split(line, " -> ")
		pStart := strings.Split(points[0], ",")
		for i := range points[1:] {
			pNext := strings.Split(points[i+1], ",")
			myMap.draw(pStart, pNext)
			pStart = pNext
		}
	}

	return myMap
}

type Point struct {
	x int
	y int
}

func (m CaveMap) print(x0, x1, y0, y1 int, sand Point) {
	for l := y0; l <= y1; l++ {
		for col := x0; col <= x1; col++ {
			if sand.x == col && sand.y == l {
				fmt.Print("o")
			} else if m[l][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (m CaveMap) simulateFallingSand() int {
	counter := 0

	maxY := len(m)
	for l := maxY - 1; l > 0; l-- {
		for x := 0; x < len(m[l]); x++ {
			if m[l][x] {
				maxY = l
				break
			}
		}
		if maxY == l {
			break
		}
	}

	sand := Point{500, 0}
	for {
		// fmt.Println(counter)
		// m.print(494, 503, 0, 9, sand)
		// fmt.Println()

		if !m[sand.y+1][sand.x] {
			if sand.y > maxY+1 {
				break
			}
			sand.y++
		} else if !m[sand.y+1][sand.x-1] {
			sand.y++
			sand.x--
		} else if !m[sand.y+1][sand.x+1] {
			sand.y++
			sand.x++
		} else {
			counter++
			m[sand.y][sand.x] = true
			sand = Point{500, 0}
		}
	}
	return counter
}

func (m CaveMap) simulateBlockSource() int {
	counter := 0

	maxY := len(m)
	for l := maxY - 1; l > 0; l-- {
		for x := 0; x < len(m[l]); x++ {
			if m[l][x] {
				maxY = l
				break
			}
		}
		if maxY == l {
			break
		}
	}

	floorY := maxY + 2

	sand := Point{500, 0}
	for {
		// fmt.Println(counter)
		// m.print(488, 512, 0, 11, sand)
		// fmt.Println()

		if !m[sand.y+1][sand.x] && sand.y+1 < floorY {
			sand.y++
		} else if !m[sand.y+1][sand.x-1] && sand.y+1 < floorY {
			sand.y++
			sand.x--
		} else if !m[sand.y+1][sand.x+1] && sand.y+1 < floorY {
			sand.y++
			sand.x++
		} else {
			counter++
			m[sand.y][sand.x] = true
			if sand.x == 500 && sand.y == 0 {
				break
			}
			sand = Point{500, 0}
		}
	}
	return counter
}

func (*Puzzle) Part1(input string) string {
	// 1. tratar o input
	// 1.1 transformar coords -> mapa
	// 2. simular areia caindo
	// 2.1 aplicar as regras físicas
	// 2.2 verificar se cai para o infinito
	myMap := parseInput(input)
	solution := myMap.simulateFallingSand()
	return fmt.Sprint(solution)
}

func (*Puzzle) Part2(input string) string {
	myMap := parseInput(input)
	solution := myMap.simulateBlockSource()
	return fmt.Sprint(solution)
}

func (*Puzzle) Notes() string {
	return "simul ampulheta: areia que cai"
}
