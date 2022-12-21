package day21

import (
	"fmt"
	"strconv"
	"strings"
)

var operators map[string]*Oper

type Puzzle struct{}

type Oper func(a, b int) int

type Monkey struct {
	name           string
	rawMonkeyValue string
	operator       string
	value          *int
	left           *Monkey
	right          *Monkey
}

func init() {
	var oSum Oper = func(a, b int) int { return a + b }
	var oSub Oper = func(a, b int) int { return a - b }
	var oMul Oper = func(a, b int) int { return a * b }
	var oDiv Oper = func(a, b int) int { return a / b }
	operators = map[string]*Oper{
		"+": &oSum,
		"-": &oSub,
		"*": &oMul,
		"/": &oDiv,
	}

}

func parseInput(raw string) map[string]*Monkey {
	monkeys := make(map[string]*Monkey, 0)
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		parts := strings.Split(line, ": ")
		monkey := Monkey{
			name:           parts[0],
			rawMonkeyValue: parts[1],
		}
		monkeys[monkey.name] = &monkey
	}
	return monkeys
}

func buildTree(monkeys map[string]*Monkey, m *Monkey) {
	if m.value != nil || m.operator != "" {
		return
	}

	if x, err := strconv.Atoi(m.rawMonkeyValue); err == nil {
		m.value = &x
		return
	}

	parts := strings.Split(m.rawMonkeyValue, " ")
	m.left = monkeys[parts[0]]
	m.right = monkeys[parts[2]]
	m.operator = parts[1]

	buildTree(monkeys, m.left)
	buildTree(monkeys, m.right)
}

func (m *Monkey) GetValue() int {
	if m.value != nil {
		return *m.value
	}

	x := (*operators[m.operator])(m.left.GetValue(), m.right.GetValue())
	m.value = &x
	return *m.value
}

func (m *Monkey) traverseToDestination(destination string, wantedValue int) (int, bool) {
	var newWanted int
	if m.name == destination {
		return wantedValue, true
	}

	if m.left != nil {
		switch m.operator {
		case "+":
			newWanted = wantedValue - *m.right.value // R = a+b | a=R-b
		case "-":
			newWanted = wantedValue + *m.right.value // R = a-b | a=R+b
		case "*":
			newWanted = wantedValue / *m.right.value // R = a*b | a=R/b
		case "/":
			newWanted = wantedValue * *m.right.value // R = a/b | a=R*b
		}

		if v, ok := m.left.traverseToDestination(destination, newWanted); ok {
			return v, true
		}
	}

	if m.right != nil {
		switch m.operator {
		case "+":
			newWanted = wantedValue - *m.left.value // R = a+b | b=R-a
		case "-":
			newWanted = *m.left.value - wantedValue // R = a-b | b=a-R
		case "*":
			newWanted = wantedValue / *m.left.value // R = a*b | b=R/a
		case "/":
			newWanted = *m.left.value / wantedValue // R = a/b | b=a/R
		}
		if v, ok := m.right.traverseToDestination(destination, newWanted); ok {
			return v, true
		}
	}

	return 0, false
}

func (*Puzzle) Part1(input string) string {
	monkeys := parseInput(input)
	buildTree(monkeys, monkeys["root"])
	return fmt.Sprint(monkeys["root"].GetValue())
}

func (*Puzzle) Part2(input string) string {
	monkeys := parseInput(input)
	buildTree(monkeys, monkeys["root"])
	monkeys["root"].GetValue() // fill all values

	if v, ok := monkeys["root"].left.traverseToDestination("humn", *monkeys["root"].right.value); ok {
		return fmt.Sprint(v)
	}

	if v, ok := monkeys["root"].right.traverseToDestination("humn", *monkeys["root"].left.value); ok {
		return fmt.Sprint(v)
	}

	panic("should not reach this state :(")
}

func (*Puzzle) Notes() string {
	return "árvore binária & backtrack"
}
