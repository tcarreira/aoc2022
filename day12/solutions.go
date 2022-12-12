package day12

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

type P struct {
	line int
	col  int
}

type PuzzleInput struct {
	mapa        [][]byte
	startPoint  P
	finishPoint P
}

func parseInput(raw string) PuzzleInput {
	puzzleInput := PuzzleInput{}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		puzzleInput.mapa = append(puzzleInput.mapa, []byte(line))
	}

	for l := range puzzleInput.mapa {
		for c, elem := range puzzleInput.mapa[l] {
			if elem == byte('S') {
				puzzleInput.startPoint = P{line: l, col: c}
				puzzleInput.mapa[l][c] = byte('a')
			} else if elem == byte('E') {
				puzzleInput.finishPoint = P{line: l, col: c}
				puzzleInput.mapa[l][c] = byte('z')
			}
		}
	}

	return puzzleInput
}

// Breadth First Search or BFS
func (p *PuzzleInput) bfs() int {
	visited := make([][]int, len(p.mapa))
	for i := range visited {
		visited[i] = make([]int, len(p.mapa[0]))
		for j := range visited[i] {
			visited[i][j] = 1 << 32
		}
	}

	visited[p.startPoint.line][p.startPoint.col] = 0
	fifo := []P{p.startPoint}
	for {
		curr := fifo[0]
		l := curr.line
		c := curr.col
		// fmt.Print("")
		// for _, line := range visited {
		// 	fmt.Println(line)
		// }
		// fmt.Println(curr, visited[l][c])

		if l == p.finishPoint.line && c == p.finishPoint.col {
			return visited[l][c]
		}

		// right
		if c+1 < len(p.mapa[l]) && visited[l][c]+1 < visited[l][c+1] && p.mapa[l][c]+1 >= p.mapa[l][c+1] {
			fifo = append(fifo, P{l, c + 1})
			visited[l][c+1] = visited[l][c] + 1
		}
		// left
		if c-1 >= 0 && visited[l][c]+1 < visited[l][c-1] && p.mapa[l][c]+1 >= p.mapa[l][c-1] {
			fifo = append(fifo, P{l, c - 1})
			visited[l][c-1] = visited[l][c] + 1
		}

		// top
		if l-1 >= 0 && visited[l][c]+1 < visited[l-1][c] && p.mapa[l][c]+1 >= p.mapa[l-1][c] {
			fifo = append(fifo, P{l - 1, c})
			visited[l-1][c] = visited[l][c] + 1
		}

		// down
		if l+1 < len(visited) && visited[l][c]+1 < visited[l+1][c] && p.mapa[l][c]+1 >= p.mapa[l+1][c] {
			fifo = append(fifo, P{l + 1, c})
			visited[l+1][c] = visited[l][c] + 1
		}

		fifo = fifo[1:]
	}
}

// Breadth First Search or BFS
func (p *PuzzleInput) bfsPart2() int {
	visited := make([][]int, len(p.mapa))
	for i := range visited {
		visited[i] = make([]int, len(p.mapa[0]))
		for j := range visited[i] {
			visited[i][j] = 1 << 32
		}
	}

	visited[p.finishPoint.line][p.finishPoint.col] = 0
	fifo := []P{p.finishPoint}
	for {
		curr := fifo[0]
		l := curr.line
		c := curr.col

		if p.mapa[l][c] == byte('a') {
			return visited[l][c]
		}

		// right
		if c+1 < len(p.mapa[l]) && visited[l][c]+1 < visited[l][c+1] && p.mapa[l][c]-1 <= p.mapa[l][c+1] {
			fifo = append(fifo, P{l, c + 1})
			visited[l][c+1] = visited[l][c] + 1
		}
		// left
		if c-1 >= 0 && visited[l][c]+1 < visited[l][c-1] && p.mapa[l][c]-1 <= p.mapa[l][c-1] {
			fifo = append(fifo, P{l, c - 1})
			visited[l][c-1] = visited[l][c] + 1
		}

		// top
		if l-1 >= 0 && visited[l][c]+1 < visited[l-1][c] && p.mapa[l][c]-1 <= p.mapa[l-1][c] {
			fifo = append(fifo, P{l - 1, c})
			visited[l-1][c] = visited[l][c] + 1
		}

		// down
		if l+1 < len(visited) && visited[l][c]+1 < visited[l+1][c] && p.mapa[l][c]-1 <= p.mapa[l+1][c] {
			fifo = append(fifo, P{l + 1, c})
			visited[l+1][c] = visited[l][c] + 1
		}

		fifo = fifo[1:]
	}
}

func (*Puzzle) Part1(input string) string {
	puzzleInput := parseInput(input)
	// fmt.Println("===================")
	// fmt.Println(byte('b'))
	// fmt.Println(byte('b') + 1)
	// fmt.Println(byte('d')+1 >= byte('a'))

	sol := puzzleInput.bfs()
	return fmt.Sprint(sol)
}

func (*Puzzle) Part2(input string) string {
	puzzleInput := parseInput(input)
	sol := puzzleInput.bfsPart2()
	return fmt.Sprint(sol)
}

func (*Puzzle) Notes() string {
	return "Caminho mais curto"
}
