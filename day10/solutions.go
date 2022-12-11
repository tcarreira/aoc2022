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
		} else if s.Instructions[i].command == "addx" {
			if s.ClockCycle+2 >= stopCycle {
				break
			}
			s.ClockCycle += 2
			s.RegisterX += s.Instructions[i].arg
		}
	}
	return s.RegisterX
}

func (*Puzzle) Part1(input string) string {
	instructions := parseInput(input)
	state := State{
		RegisterX:    1,
		ClockCycle:   0,
		Instructions: instructions,
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
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
