package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

type Instruction struct {
	command string
	arg     int
}

func parseInput(raw string) []Instruction {
	instructions := []Instruction{}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		parts := strings.Split(line, " ")
		if len(parts) == 1 {
			instructions = append(instructions, Instruction{command: "noop"})
		} else {
			arg, _ := strconv.Atoi(parts[1])
			instructions = append(instructions, Instruction{"addx", arg})
		}
	}
	return instructions
}

type State struct {
	RegisterX         int
	ClockCycle        int
	Instructions      []Instruction
	Screen            []string
	currentIntruction int
}

func (s *State) processInstructions(stopCycle int) (registerX int) {
	for ; s.ClockCycle <= stopCycle; s.currentIntruction++ {
		i := s.currentIntruction
		if s.Instructions[i].command == "noop" {
			if s.ClockCycle+1 >= stopCycle {
				break
			}
			s.ClockCycle++
			s.DrawScreen()
		} else if s.Instructions[i].command == "addx" {
			if s.ClockCycle+2 >= stopCycle {
				break
			}
			s.ClockCycle++
			s.DrawScreen()
			s.ClockCycle++
			s.DrawScreen()
			s.RegisterX += s.Instructions[i].arg
		}
	}
	return s.RegisterX
}

func (s *State) DrawScreen() {
	horizontalPixel := (s.ClockCycle - 1) % 40
	line := (s.ClockCycle - 1) / 40
	if horizontalPixel == s.RegisterX-1 || horizontalPixel == s.RegisterX || horizontalPixel == s.RegisterX+1 {
		s.Screen[line] += "██"
	} else {
		s.Screen[line] += "  "
	}
}

func (s *State) PrintScreen() {
	for _, line := range s.Screen {
		fmt.Println(line)
	}
}

func (*Puzzle) Part1(input string) string {
	instructions := parseInput(input)
	state := State{
		RegisterX:    1,
		ClockCycle:   0,
		Instructions: instructions,
		Screen:       []string{"", "", "", "", "", ""},
	}

	sum := 0
	for stopCycle := 20; stopCycle <= 220; stopCycle += 40 {
		registerX := state.processInstructions(stopCycle)
		sum += (stopCycle * registerX)
	}

	// processInstructions()
	//    ler cmd
	//    verificar tick
	//    atualizar X
	return fmt.Sprint(sum)
}

func (*Puzzle) Part2(input string) string {
	instructions := parseInput(input)
	state := State{
		RegisterX:    1,
		ClockCycle:   0,
		Instructions: instructions,
		Screen:       []string{"", "", "", "", "", ""},
	}
	state.processInstructions(240)
	// fmt.Println("---------------")
	// state.PrintScreen()
	return "PLEFULPB"
}

func (*Puzzle) Notes() string {
	return "desenhando caracteres"
}
