package day22

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	maxCol   int
	maxLine  int
	cubeSize int
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

func (mm *MonkeyMap) stepOnCube(p *P, cubeWrapper CubeWrapper) bool {
	// test for wrapping around
	if newP, ok := cubeWrapper[*p]; ok {
		if (*mm)[P{line: newP.line, col: newP.col}].isWall {
			return false
		}
		p.line = newP.line
		p.col = newP.col
		p.dir = newP.dir
		return true
	}

	switch p.dir {
	case 0: // face right
		newCol := p.col + 1
		tile, ok := (*mm)[P{line: p.line, col: newCol}]
		if !ok { //wrap around
			panic("should not be reachable")
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
			panic("should not be reachable")
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
			panic("should not be reachable")
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
			panic("should not be reachable")
		}
		if tile.isWall {
			return false
		}
		p.line = newLine
		return true
	}
	panic(fmt.Errorf("not reachable, %T , %#v", p, p))
}

func (mm *MonkeyMap) moveOnCube(pos *P, instr string, cubeWrapper CubeWrapper) {
	if steps, err := strconv.Atoi(instr); err == nil {
		for i := 0; i < steps; i++ {
			if !mm.stepOnCube(pos, cubeWrapper) {
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

func shrinkTo1x1Cube(mm MonkeyMap) string {
	cube1x1 := ""

	F := byte('A')
	// example is 4x3, real input is 4x3
	for l := 0; l < 4; l++ {
		for c := 0; c < 4; c++ {
			if _, ok := mm[P{line: l * cubeSize, col: c * cubeSize}]; ok {
				cube1x1 += string(F)
				F++
			} else {
				cube1x1 += "."
			}
		}
		cube1x1 += "\n"
	}
	// fmt.Println(cube1x1)
	return cube1x1
}

type CubeWrapper map[P]P

func buildCubeWrapper(mm MonkeyMap) CubeWrapper {
	cubeWrapper := CubeWrapper{}

	cube1x1 := shrinkTo1x1Cube(mm)
	if cube1x1 == "..A.\nBCD.\n..EF\n....\n" { // example
		// ..A.
		// BCD.
		// ..EF

		// 14 transitions between faces

		// A(0,2) --up--> B(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 0 + 0*cubeSize, col: x + 2*cubeSize, dir: 3}
			newP := P{line: 0 + 1*cubeSize, col: x + 0*cubeSize, dir: 1}
			cubeWrapper[oldP] = newP
		}

		// A(0,2) --right--> F(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: x + 0*cubeSize, col: 3*cubeSize - 1, dir: 0}
			newP := P{line: 3*cubeSize - 1 - x, col: 4*cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// D(1,2) --right--> F(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: x + 1*cubeSize, col: 3*cubeSize - 1, dir: 0}
			newP := P{line: 2 * cubeSize, col: 4*cubeSize - 1 - x, dir: 1}
			cubeWrapper[oldP] = newP
		}

		// F(2,3) --up-----> D(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2 * cubeSize, col: x + 3*cubeSize, dir: 3}
			newP := P{line: 2*cubeSize - 1 - x, col: 3*cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// F(2,3) --right--> A(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: x + 2*cubeSize, col: 4*cubeSize - 1, dir: 0}
			newP := P{line: 1*cubeSize - 1 - x, col: 3*cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// F(2,3) --down---> B(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 3*cubeSize - 1, col: 3*cubeSize + x, dir: 1}
			newP := P{line: 2*cubeSize - 1 - x, col: 0, dir: 0}
			cubeWrapper[oldP] = newP
		}

		// E(2,2) --down---> B(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 3*cubeSize - 1, col: 2*cubeSize + x, dir: 1}
			newP := P{line: 2*cubeSize - 1, col: 1*cubeSize - 1 - x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// E(2,2) --left---> C(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2*cubeSize + x, col: 2 * cubeSize, dir: 2}
			newP := P{line: 2*cubeSize - 1, col: 2*cubeSize - 1 - x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// C(1,1) --down---> E(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2*cubeSize - 1, col: 2*cubeSize - 1 - x, dir: 1}
			newP := P{line: 3*cubeSize - 1 - x, col: 2 * cubeSize, dir: 0}
			cubeWrapper[oldP] = newP
		}

		// B(1,0) --down---> E(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2*cubeSize - 1, col: x, dir: 1}
			newP := P{line: 3*cubeSize - 1, col: 3*cubeSize - 1 - x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// B(1,0) --left---> F(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 1*cubeSize + x, col: 0, dir: 2}
			newP := P{line: 3*cubeSize - 1, col: 4*cubeSize - 1 - x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// B(1,0) --up-----> A(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 1 * cubeSize, col: x, dir: 3}
			newP := P{line: 0, col: 3*cubeSize - 1 - x, dir: 1}
			cubeWrapper[oldP] = newP
		}

		// C(1,1) --up-----> A(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 1 * cubeSize, col: 1*cubeSize + x, dir: 3}
			newP := P{line: x, col: 2 * cubeSize, dir: 0}
			cubeWrapper[oldP] = newP
		}

		// A(0,2) --left---> C(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: x, col: 2 * cubeSize, dir: 2}
			newP := P{line: 1 * cubeSize, col: 2*cubeSize - 1 - x, dir: 1}
			cubeWrapper[oldP] = newP
		}

	} else if cube1x1 == ".AB.\n.C..\nDE..\nF...\n" { // my real input
		// .AB.
		// .C..
		// DE..
		// F...

		// 14 transitions between faces

		// A(0,1) --up-----> F(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 0, col: cubeSize + x, dir: 3}
			newP := P{line: 3*cubeSize + x, col: 0, dir: 0}
			cubeWrapper[oldP] = newP
		}

		// B(0,2) --up-----> F(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 0, col: 2*cubeSize + x, dir: 3}
			newP := P{line: 4*cubeSize - 1, col: x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// B(0,2) --right--> E(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: x, col: 3*cubeSize - 1, dir: 0}
			newP := P{line: 3*cubeSize - 1 - x, col: 2*cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// B(0,2) --down---> C(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: cubeSize - 1, col: 2*cubeSize + x, dir: 1}
			newP := P{line: cubeSize + x, col: 2*cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// C(1,1) --right--> B(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: cubeSize + x, col: 2*cubeSize - 1, dir: 0}
			newP := P{line: cubeSize - 1, col: 2*cubeSize + x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// E(2,1) --right--> B(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2*cubeSize + x, col: 2*cubeSize - 1, dir: 0}
			newP := P{line: cubeSize - 1 - x, col: 3*cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// E(2,1) --down---> F(<)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 3*cubeSize - 1, col: cubeSize + x, dir: 1}
			newP := P{line: 3*cubeSize + x, col: cubeSize - 1, dir: 2}
			cubeWrapper[oldP] = newP
		}

		// F(3,0) --right--> E(^)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 3*cubeSize + x, col: cubeSize - 1, dir: 0}
			newP := P{line: 3*cubeSize - 1, col: cubeSize + x, dir: 3}
			cubeWrapper[oldP] = newP
		}

		// F(3,0) --down---> B(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 4*cubeSize - 1, col: x, dir: 1}
			newP := P{line: 0, col: 2*cubeSize + x, dir: 1}
			cubeWrapper[oldP] = newP
		}

		// F(3,0) --left---> A(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 3*cubeSize + x, col: 0, dir: 2}
			newP := P{line: 0, col: cubeSize + x, dir: 1}
			cubeWrapper[oldP] = newP
		}

		// D(2,0) --left---> A(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2*cubeSize + x, col: 0, dir: 2}
			newP := P{line: cubeSize - 1 - x, col: cubeSize, dir: 0}
			cubeWrapper[oldP] = newP
		}

		// D(2,0) --up-----> C(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: 2 * cubeSize, col: x, dir: 3}
			newP := P{line: cubeSize + x, col: cubeSize, dir: 0}
			cubeWrapper[oldP] = newP
		}

		// C(1,1) --left---> D(v)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: cubeSize + x, col: cubeSize, dir: 2}
			newP := P{line: 2 * cubeSize, col: x, dir: 1}
			cubeWrapper[oldP] = newP
		}

		// A(0,1) --left---> D(>)
		for x := 0; x < cubeSize; x++ {
			oldP := P{line: x, col: cubeSize, dir: 2}
			newP := P{line: 3*cubeSize - 1 - x, col: 0, dir: 0}
			cubeWrapper[oldP] = newP
		}

	} else {
		fmt.Printf("Found cube planification: %q\n", cube1x1)
		panic("unknown cube planification (half-hardecoded-solution)")
	}

	_ = cube1x1
	return cubeWrapper
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
	monkeyMap, instructions := parseInput(input)
	position := findInitialPosition(monkeyMap)

	if maxLine < 30 {
		cubeSize = 4 // example
	} else {
		cubeSize = 50 // real input
	}
	cubeWrapper := buildCubeWrapper(monkeyMap)

	for _, instruction := range instructions {
		monkeyMap.moveOnCube(&position, instruction, cubeWrapper)
	}
	return fmt.Sprint((position.line+1)*1000 + (position.col+1)*4 + position.dir)
}

func (*Puzzle) Notes() string {
	return "caminhar na face do cubo (planif hardc)"
}
