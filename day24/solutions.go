package day24

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

const (
	Right = '>'
	Left  = '<'
	Up    = '^'
	Down  = 'v'
)

var (
	dx, dX, dy, dY int
)

type Blizzard struct {
	x, y int
	dir  rune
}
type Blizards []Blizzard

func parseInput(raw string) Blizards {
	blizs := Blizards{}

	lines := strings.Split(strings.TrimSpace(raw), "\n")
	dY = len(lines) - 1
	for y, line := range lines {
		dX = len(line) - 1
		for x, c := range line {
			switch c {
			case Right, Left, Up, Down:
				blizs = append(blizs, Blizzard{x, y, c})
			}
		}
	}
	return blizs
}

type P struct{ X, Y int }
type P3 struct{ X, Y, t int }

type BMap [][]int

func wrap(n, nmin, nmax int) int {
	mod := (nmax + 1) - nmin
	if (n - nmin) < 0 {
		return nmax + (n+1-nmin)%mod
	}
	return (n-nmin)%mod + nmin
}

func (bs *Blizards) Step(n int) BMap {
	bMap := make(BMap, dY+1)
	for y := range bMap {
		bMap[y] = make([]int, dX+1)
	}

	for _, b := range *bs {
		switch b.dir {
		case Right:
			bMap[b.y][wrap(b.x+n, dx+1, dX-1)]++
		case Left:
			bMap[b.y][wrap(b.x-n, dx+1, dX-1)]++
		case Up:
			bMap[wrap(b.y-n, dy+1, dY-1)][b.x]++
		case Down:
			bMap[wrap(b.y+n, dy+1, dY-1)][b.x]++
		}
	}
	return bMap
}

func (bmap *BMap) isClear(p P3) bool {
	if (p.X == dx+1 && p.Y == dy) || (p.Y == dY && p.X == dX-1) {
		return true
	}
	if p.X <= dx || p.X >= dX || p.Y <= dy || p.Y >= dY {
		return false
	}
	if (*bmap)[p.Y][p.X] > 0 { // has blizzard
		return false
	}
	return true
}

func (p P3) nextStepPoints(t int) []P3 {
	return []P3{
		{p.X, p.Y, t},     // stay
		{p.X + 1, p.Y, t}, // right
		{p.X - 1, p.Y, t}, // left
		{p.X, p.Y + 1, t}, // down
		{p.X, p.Y - 1, t}, // up
	}
}

func (*Puzzle) Part1(input string) string {
	blizs := parseInput(input)

	// setup State
	startPoint, finishPoint := P3{dx + 1, dy, 0}, P{dX - 1, dY}

	// main cycle
	queue := map[P3]bool{startPoint: true}
	for t := 1; ; t++ {
		bmap := blizs.Step(t)
		nextQueue := map[P3]bool{}

		for p := range queue {
			if p.X == finishPoint.X && p.Y == finishPoint.Y {
				return fmt.Sprint(t - 1)
			}

			for _, nextP := range p.nextStepPoints(t) {
				if _, ok := nextQueue[nextP]; ok {
					continue
				}
				if !bmap.isClear(nextP) {
					continue
				}
				nextQueue[nextP] = true
			}
			queue = nextQueue
		}
	}
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
