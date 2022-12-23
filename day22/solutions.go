package day22

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	maxCol  int
	maxLine int
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

type MonkeyMap map[P]Tile

func parseInput(raw string) (MonkeyMap, []string) {
	monkeyMap := MonkeyMap{}

	parts := strings.Split(raw, "\n\n")
	mapStrLines := strings.Split(parts[0], "\n")
	maxLine = len(mapStrLines)
	for line, lineStr := range mapStrLines {
		for col, c := range lineStr {
			if len(lineStr) > maxCol {
				maxCol = col
			}

			switch c {
			case '#':
				monkeyMap[P{line: line, col: col}] = Tile{isWall: true}
			case '.':
				monkeyMap[P{line: line, col: col}] = Tile{isOpen: true}
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

func findInitialPosition(monkeyMap MonkeyMap) P {
	for i := 0; ; i++ {
		if _, ok := monkeyMap[P{line: 0, col: i}]; ok {
			return P{line: 0, col: i, dir: 0}
		}
	}
	panic("not reachable")
}

func (mm *MonkeyMap) step(p *P) bool {
	switch p.dir {
	case 0: // face right
		newCol := p.col + 1
		tile, ok := (*mm)[P{line: p.line, col: newCol}]
		if !ok { //wrap around
			for i := 0; ; i++ {
				if tile, ok = (*mm)[P{line: p.line, col: i}]; ok {
					newCol = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.col = newCol
		return true

	case 1: // face down
		newLine := p.line + 1
		tile, ok := (*mm)[P{line: newLine, col: p.col}]
		if !ok { //wrap around
			for i := 0; ; i++ {
				if tile, ok = (*mm)[P{line: i, col: p.col}]; ok {
					newLine = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.line = newLine
		return true

	case 2: // face left
		newCol := p.col - 1
		tile, ok := (*mm)[P{line: p.line, col: newCol}]
		if !ok { //wrap around
			for i := maxCol; ; i-- {
				if tile, ok = (*mm)[P{line: p.line, col: i}]; ok {
					newCol = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.col = newCol
		return true

	case 3: // face up
		newLine := p.line - 1
		tile, ok := (*mm)[P{line: newLine, col: p.col}]
		if !ok { //wrap around
			for i := maxLine; ; i-- {
				if tile, ok = (*mm)[P{line: i, col: p.col}]; ok {
					newLine = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.line = newLine
		return true
	}
	panic(fmt.Errorf("not reachable, %T , %#v", p, p))
}

func (mm *MonkeyMap) move(pos *P, instr string) {
	// fmt.Println(pos, instr)
	if steps, err := strconv.Atoi(instr); err == nil {
		for i := 0; i < steps; i++ {
			if !mm.step(pos) {
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
			pos.dir = 3
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
	return fmt.Sprint((position.line+1)*1000 + (position.col+1)*4 + position.dir)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
