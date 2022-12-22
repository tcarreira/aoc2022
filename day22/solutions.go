package day22

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Puzzle struct{}

type P struct {
	line int
	col  int
	dir  int
}

type Tile struct {
	isOpen bool
	isWall bool
}

func parseInput(raw string) (MonkeyMap, []string) {
	var monkeyMap MonkeyMap

	parts := strings.Split(raw, "\n\n")
	for l, line := range strings.Split(parts[0], "\n") {
		monkeyMap = append(monkeyMap, make([]Tile, len(line)))
		for col, c := range line {
			switch c {
			case '#':
				monkeyMap[l][col].isWall = true
			case '.':
				monkeyMap[l][col].isOpen = true
			}
		}
	}

	instructions := []string{}

	instructionStr := strings.TrimSpace(parts[1])
	for i := 0; i < len(instructionStr); i++ {
		c := rune(parts[1][i])
		if unicode.IsLetter(c) {
			instructions = append(instructions, string(c))
		} else {
			j := i
			for _, c2 := range instructionStr[i:] {
				if unicode.IsLetter(c2) {
					break
				}
				j++ //4
			}
			instructions = append(instructions, parts[1][i:j])
			i = j - 1
		}
	}

	return monkeyMap, instructions
}

type MonkeyMap [][]Tile

func findInitialPosition(monkeyMap MonkeyMap) P {
	for i, t := range monkeyMap[0] {
		if t.isOpen {
			return P{line: 0, col: i, dir: 0}
		}
	}
	panic("not reachable")
}

func (mm *MonkeyMap) step(p *P) bool {
	switch p.dir {
	case 0: // face right
		if p.col+1 >= len((*mm)[p.line]) || (!(*mm)[p.line][j].isWall) { // wrap around
			for j := 0; ; j++ {
				if (*mm)[p.line][j].isOpen {
					p.col = j
					return true
				}
				if (*mm)[p.line][j].isWall {
					return false
				}
			}
		}
		if (*mm)[p.line][p.col+1].isWall { // wall
			return false
		}
		p.col++
	case 1: // face down
		if p.line+1 >= len((*mm)) { // wrap around
			for j := 0; ; j++ {
				if len((*mm)[j]) <= p.col {
					continue
				}
				if (*mm)[j][p.col].isOpen {
					p.line = j
					return true
				}
				if (*mm)[j][p.col].isWall {
					return false
				}
			}
		}
		if (*mm)[p.line][p.col+1].isWall { // wall
			return false
		}
		p.col++

	case 2: // face left
		if p.col-1 <= 0 { // wrap around
			for j := 0; ; j++ {
				if (*mm)[p.line][j].isOpen {
					p.col = j
					return true
				}
				if (*mm)[p.line][j].isWall {
					return false
				}
			}
		}
		if (*mm)[p.line][p.col+1].isWall { // wall
			return false
		}
		p.col++
	}
	panic("not reachable")
}

func (mm *MonkeyMap) move(pos *P, instr string) {
	if steps, err := strconv.Atoi(instr); err == nil {
		for i := 0; i < steps; i++ {
			if !pos.step(pos) {
				return
			}
		}

		return
	}

	// instr is letter
	switch instr {
	case "R":
		pos.dir = (pos.dir + 1) % 4
	case "L":
		if pos.dir == 0 {
			pos.dir = 4
		} else {
			pos.dir--
		}
	}

}

func (*Puzzle) Part1(input string) string {
	monkeyMap, instructions := parseInput(input)
	position := findInitialPosition(monkeyMap)

	for _, instruction := range instructions {
		monkeyMap.move(&position, instruction)

	}
	fmt.Println(position, instructions)
	// buildMap
	// simulateWalk
	//   colisionWall
	//   endMap -> wrapAround
	return "-"
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
