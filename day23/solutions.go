package day23

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

type P struct {
	x int
	y int
}

type Elfs map[P]*P

func parseInput(raw string) Elfs {
	elfsMap := make(Elfs)
	for y, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		for x, c := range line {
			if c == '#' {
				elfsMap[P{x, y}] = nil
			}
		}
	}
	return elfsMap
}

func (me *Elfs) Plan(round int) map[P]int {
	directions := []string{"N", "S", "W", "E"}
	planCounter := make(map[P]int)

	for elf := range *me {

		adjFound := false
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				if _, ok := (*me)[P{elf.x + x, elf.y + y}]; ok {
					adjFound = true
					break
				}
			}
		}
		if !adjFound {
			continue
		}

		for dirIdx := 0; dirIdx < 4; dirIdx++ {
			dir := directions[(round+dirIdx)%4]

			checkPos := [3]P{}
			newP := P{}
			switch dir {
			case "N":
				checkPos[0] = P{elf.x - 1, elf.y - 1}
				checkPos[1] = P{elf.x, elf.y - 1}
				checkPos[2] = P{elf.x + 1, elf.y - 1}
				newP = P{elf.x, elf.y - 1}
			case "S":
				checkPos[0] = P{elf.x - 1, elf.y + 1}
				checkPos[1] = P{elf.x, elf.y + 1}
				checkPos[2] = P{elf.x + 1, elf.y + 1}
				newP = P{elf.x, elf.y + 1}
			case "W":
				checkPos[0] = P{elf.x - 1, elf.y - 1}
				checkPos[1] = P{elf.x - 1, elf.y}
				checkPos[2] = P{elf.x - 1, elf.y + 1}
				newP = P{elf.x - 1, elf.y}
			case "E":
				checkPos[0] = P{elf.x + 1, elf.y - 1}
				checkPos[1] = P{elf.x + 1, elf.y}
				checkPos[2] = P{elf.x + 1, elf.y + 1}
				newP = P{elf.x + 1, elf.y}
			}

			_, ok0 := (*me)[checkPos[0]]
			_, ok1 := (*me)[checkPos[1]]
			_, ok2 := (*me)[checkPos[2]]
			if !ok0 && !ok1 && !ok2 {
				// update elf next planed position
				(*me)[elf] = &newP

				// update aux count
				planCounter[newP]++
				break
			}
		}
	}
	return planCounter
}

func (me *Elfs) ExecutePlan(planCounter map[P]int) Elfs {
	newElfsMap := make(Elfs)
	for elf, plan := range *me {
		if plan != nil && planCounter[*plan] == 1 {
			newElfsMap[*plan] = nil
		} else {
			newElfsMap[elf] = nil
		}
	}
	return newElfsMap
}

func (me *Elfs) Boundary() (x, X, y, Y int) {
	for elf := range *me {
		if elf.x < x {
			x = elf.x
		}
		if elf.x > X {
			X = elf.x
		}
		if elf.y < y {
			y = elf.y
		}
		if elf.y > Y {
			Y = elf.y
		}
	}
	return
}

func (me *Elfs) Print() {
	x, X, y, Y := (*me).Boundary()

	for _y := y; _y <= Y; _y++ {
		for _x := x; _x <= X; _x++ {
			if _, ok := (*me)[P{_x, _y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (me *Elfs) CountEmtpyTiles() int {
	x, X, y, Y := (*me).Boundary()

	elfCount := len(*me)
	fullRect := (X - x + 1) * (Y - y + 1)

	return fullRect - elfCount
}

func (*Puzzle) Part1(input string) string {

	elfs := parseInput(input)

	// elfs.Print()
	for i := 0; i < 10; i++ {
		planCounter := elfs.Plan(i)
		elfs = elfs.ExecutePlan(planCounter)
		// elfs.Print()
	}

	// 1 - planning ::: map[Pplan]int
	// 1.1. no adjecent: do not move
	// 1.2. verificar direção1, dir2, dir3, dir4: move dir  ::: []str{"N","S","W","E"}

	// 2 - exec
	// 2.1. move if no other elf planed the same

	return fmt.Sprint(elfs.CountEmtpyTiles())
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
