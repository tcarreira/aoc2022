package day17

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

type Piece struct {
	width int
	shape [4][4]bool
}

func buildPieces() []Piece {
	return []Piece{
		{
			width: 4,
			shape: [4][4]bool{
				{true, true, true, true},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
		},
		{
			width: 3,
			shape: [4][4]bool{
				{false, true, false, false},
				{true, true, true, false},
				{false, true, false, false},
				{false, false, false, false},
			},
		},
		{
			width: 3,
			shape: [4][4]bool{
				{true, true, true, false},
				{false, false, true, false},
				{false, false, true, false},
				{false, false, false, false},
			},
		},
		{
			width: 1,
			shape: [4][4]bool{
				{true, false, false, false},
				{true, false, false, false},
				{true, false, false, false},
				{true, false, false, false},
			},
		},
		{
			width: 2,
			shape: [4][4]bool{
				{true, true, false, false},
				{true, true, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
		},
	}
}

// [7]bool

type Cache struct {
	height     int
	pieceCount int
}

type Chamber struct {
	pieceCounter int
	moveIdx      int
	height       int
	moves        string
	pieces       []Piece
	state        [][7]bool
	cache        map[string]Cache
}

func (c *Chamber) isColision(pos [2]int, piece Piece) bool {
	if pos[0] < 0 || pos[1] < 0 || pos[1]+piece.width > len(c.state[0]) {
		return true
	}
	for line := 0; line < len(piece.shape); line++ {
		for col := 0; col < len(piece.shape[line]); col++ {
			if piece.shape[line][col] && c.state[pos[0]+line][pos[1]+col] {
				return true
			}
		}
	}
	return false
}
func (c *Chamber) savePiece(pos [2]int, piece Piece) {
	for line := 0; line < len(piece.shape); line++ {
		for col := 0; col < len(piece.shape[line]); col++ {
			if piece.shape[line][col] {
				c.state[pos[0]+line][pos[1]+col] = true

				// update c.height
				if pos[0]+line+1 > c.height {
					c.height = pos[0] + line + 1
				}
			}
		}
	}
}

func (c *Chamber) printState(pos [2]int, piece Piece) {
	for line := len(c.state) - 1; line >= 0; line-- {
		tmpLine := ""
		for col := 0; col < len(c.state[line]); col++ {
			if c.state[line][col] {
				tmpLine += "#"
			} else {
				tmpLine += "."
			}
		}
		fmt.Println(tmpLine)
	}
	fmt.Println()
}

func (c *Chamber) NextPiece() {
	c.pieceCounter++
	pos := [2]int{c.height + 3, 2}
	piece := c.pieces[(c.pieceCounter-1)%len(c.pieces)]

	newLines := (c.height + 7) - len(c.state)
	for i := 0; i < newLines; i++ {
		c.state = append(c.state, [7]bool{})
	}

	// c.printState(pos, piece)
	for { // while falling piece
		// Move horizontal
		newPos := [2]int{pos[0], pos[1]}
		switch c.moves[c.moveIdx] {
		case '>':
			newPos[1]++
		case '<':
			newPos[1]--
		}

		if !c.isColision(newPos, piece) {
			pos = newPos
		}

		c.moveIdx = (c.moveIdx + 1) % len(c.moves)

		// move down
		if c.isColision([2]int{pos[0] - 1, pos[1]}, piece) {
			// save state
			c.savePiece(pos, piece)
			return
		}
		pos[0]--

	}

	// moveHorizontal()
	// checkColision() -> revert
	// moveDown()
	// checkColision() -> revert, return
}

func (c *Chamber) updateSeenState() Cache {
	if len(c.state) < 30 {
		return Cache{}
	}

	pieceIdx := c.pieceCounter % len(c.pieces)
	stateHash := fmt.Sprintf("%03d:%05d:", pieceIdx, c.moveIdx)
	for l := 0; l < 30; l++ {
		lineHash := 0
		for x := 0; x < 7; x++ {
			if c.state[l][6-x] {
				lineHash |= 1 << x
			}
		}
		stateHash += fmt.Sprintf("%03d", lineHash)
	}

	if x, ok := c.cache[stateHash]; ok {
		return x
	}
	c.cache[stateHash] = Cache{height: c.height, pieceCount: c.pieceCounter}
	return Cache{}
}

// 8: f f f f f f f
// 7: f f f f f f f
// 6: f f f f f f f
// 5: f f f f f f f
// 4: f f f f f f f
// 3: f f f f f f f
// 2: f f f f f f f
// 1: f f f f f f f
// 0: f f f f f f f
//    0 1 2 3 4 5 6

// parseInput()
// encodePieces()  ==== []Piece
// simulador
//
//	estado (ocupação da câmara)
//	simul peca a cair
func (*Puzzle) Part1(input string) string {
	chamber := Chamber{
		moves:  strings.TrimSpace(input),
		pieces: buildPieces(),
		state:  [][7]bool{},
	}

	for i := 0; i < 2022; i++ {
		chamber.NextPiece()
	}

	return fmt.Sprint(chamber.height)
}

func (*Puzzle) Part2(input string) string {
	chamber := Chamber{
		moves:  strings.TrimSpace(input),
		pieces: buildPieces(),
		state:  [][7]bool{},
		cache:  make(map[string]Cache),
	}

	// 0123456789012345678901234567890123456789012345678901234567890123456 T
	// uuRRRRRRR.......................................................xxx
	// unique - REPEATED - ...notCalculated... - rest
	for i := 0; i < 1000000000000; i++ {
		chamber.NextPiece()

		cacheState := chamber.updateSeenState() // 1
		// c.current = 8
		if cacheState.height > 0 {
			remaining := 1000000000000 - chamber.pieceCounter
			repeatedHeight := chamber.height - cacheState.height
			repeatedLen := chamber.pieceCounter - cacheState.pieceCount
			repeatedTotalHeight := repeatedHeight * (remaining / repeatedLen)

			// chamber.height += repeatedTotalHeight
			// i = remaining
			for j := 0; j < remaining%repeatedLen; j++ {
				chamber.NextPiece()
			}
			chamber.height += repeatedTotalHeight
			break
		}
	}

	return fmt.Sprint(chamber.height)
}

func (*Puzzle) Notes() string {
	return ""
}
