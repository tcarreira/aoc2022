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
	value          *int
	operator       *Oper
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
	if m.value != nil || m.operator != nil {
		return
	}

	if x, err := strconv.Atoi(m.rawMonkeyValue); err == nil {
		m.value = &x
		return
	}

	parts := strings.Split(m.rawMonkeyValue, " ")
	m.left = monkeys[parts[0]]
	m.right = monkeys[parts[2]]
	m.operator = operators[parts[1]]

	buildTree(monkeys, m.left)
	buildTree(monkeys, m.right)
}

func (m *Monkey) GetValue() int {
	if m.value != nil {
		return *m.value
	}

	return (*m.operator)(m.left.GetValue(), m.right.GetValue())
}

func (*Puzzle) Part1(input string) string {
	monkeys := parseInput(input)
	// fmt.Println(monkeys)
	buildTree(monkeys, monkeys["root"])
	// fmt.Println(len(monkeys))
	// for _, m := range monkeys {
	// 	fmt.Printf("%p, %+v\n", &m, m)
	// }
	return fmt.Sprint(monkeys["root"].GetValue())
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
