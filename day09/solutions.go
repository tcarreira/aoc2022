package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

type Move struct {
	direction string
	steps     int
}

func parseInput(raw string) []Move {
	moves := []Move{}

	lines := strings.Split(strings.TrimSpace(raw), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		moves = append(moves, Move{
			direction: parts[0],
			steps:     steps,
		})
	}
	return moves
}

// []Pos(L,C) -> [ (0,0), (0,1), (0,2), (1,2), (2,2), (2,1), (2,0), (2,-1) ]

type Pos struct {
	line int
	col  int
}

func (p *Pos) moveHead(mov Move) {
	switch mov.direction {
	case "R":
		p.col++
	case "L":
		p.col--
	case "U":
		p.line--
	case "D":
		p.line++
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// . . . .
// . . T .
// . . . .
// . H . .

func (t *Pos) moveTail(h Pos) {
	//  T      H
	// 1,2 :: 3,1  -> 1,1 ou -1,-1
	lineDiff := t.line - h.line
	colDiff := t.col - h.col

	if abs(colDiff) <= 1 && abs(lineDiff) <= 1 {
		return
	}

	if lineDiff == 0 {
		if colDiff > 0 {
			t.col--
		} else {
			t.col++
		}
	} else if colDiff == 0 {
		if lineDiff > 0 {
			t.line--
		} else {
			t.line++
		}
	} else {
		if colDiff > 0 && lineDiff > 0 {
			t.line--
			t.col--
		} else if colDiff > 0 && lineDiff < 0 {
			t.line++
			t.col--
		} else if colDiff < 0 && lineDiff > 0 {
			t.line--
			t.col++
		} else if colDiff < 0 && lineDiff < 0 {
			t.line++
			t.col++
		} else {
			panic("ahhhhhh")
		}
	}
}

type TailTrail map[Pos]struct{}

func (t TailTrail) updateTailTrail(tail Pos) {
	if _, ok := t[tail]; !ok {
		t[tail] = struct{}{}
	}
}

func simulator(moves []Move) int {
	head := Pos{0, 0}
	tail := Pos{0, 0}
	tailTrail := TailTrail{Pos{0, 0}: struct{}{}}
	// tailTrail[Pos{0, 0}] = struct{}{} // isto é o mesmo que está acima

	for _, mov := range moves {
		for step := 0; step < mov.steps; step++ {
			head.moveHead(mov)
			tail.moveTail(head)
			tailTrail.updateTailTrail(tail)
		}
	}

	return len(tailTrail)
}

func simulatorPart2(moves []Move) int {
	knots := []Pos{}
	for i := 0; i < 10; i++ {
		knots = append(knots, Pos{0, 0})
	}

	tailTrail := TailTrail{Pos{0, 0}: struct{}{}}

	for _, mov := range moves {
		for step := 0; step < mov.steps; step++ {
			knots[0].moveHead(mov)
			for i := range knots[1:] {
				knots[i+1].moveTail(knots[i]) // move next, based on previous
			}
			tailTrail.updateTailTrail(knots[9])
		}
	}
	return len(tailTrail)
}

func (*Puzzle) Part1(input string) string {
	// simularMovimentos()
	//   headPos.move(mov)
	//   tailPos.move(headPos)

	// [][]bool
	// contarPosicoesCauda()
	moves := parseInput(input)
	solution := simulator(moves)
	return fmt.Sprint(solution)
}

func (*Puzzle) Part2(input string) string {
	moves := parseInput(input)
	solution := simulatorPart2(moves)
	return fmt.Sprint(solution)
}

func (*Puzzle) Notes() string {
	return "map/dict/set coord visitadas"
}
