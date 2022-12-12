package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct{}

type Monkey struct {
	items              []int
	itemsCount         int
	operation          Operation
	testDivisible      int
	throwMonkeyIfTrue  int
	throwMonkeyIfFalse int
}

type Operation struct {
	oper string
	arg  string
}

type MonkeyState struct {
	monkeys []Monkey
}

func parseInput(raw string) MonkeyState {
	monkeyState := MonkeyState{}
	for mID, monkey := range strings.Split(strings.TrimSpace(raw), "\n\n") {
		monkeyState.monkeys = append(monkeyState.monkeys, Monkey{})
		lines := strings.Split(monkey, "\n")

		// Starting items:
		startingItemsParts := strings.Split(lines[1], "Starting items: ")
		itemsParts := strings.Split(startingItemsParts[1], ", ")
		for _, item := range itemsParts {
			worryLevel, _ := strconv.Atoi(item)
			monkeyState.monkeys[mID].items = append(monkeyState.monkeys[mID].items, worryLevel)
		}

		// Operation:
		operationParts := strings.Split(lines[2], "Operation: ")
		parts := strings.Split(operationParts[1], " ")
		monkeyState.monkeys[mID].operation = Operation{
			oper: parts[3],
			arg:  parts[4],
		}

		// Test:
		parts = strings.Split(lines[3], " ")
		divisibleNum, _ := strconv.Atoi(parts[len(parts)-1])
		monkeyState.monkeys[mID].testDivisible = divisibleNum

		// If true:
		parts = strings.Split(lines[4], " ")
		monkeyID, _ := strconv.Atoi(parts[len(parts)-1])
		monkeyState.monkeys[mID].throwMonkeyIfTrue = monkeyID

		// If false:
		parts = strings.Split(lines[5], " ")
		monkeyID, _ = strconv.Atoi(parts[len(parts)-1])
		monkeyState.monkeys[mID].throwMonkeyIfFalse = monkeyID
	}

	return monkeyState
}

func (ms *MonkeyState) processRound() {
	for monkeyID, monkey := range ms.monkeys {
		ms.monkeys[monkeyID].itemsCount += len(monkey.items)
		for _, worryLevel := range monkey.items {
			var err error

			// part 1: operation
			arg := worryLevel
			if monkey.operation.arg != "old" {
				arg, err = strconv.Atoi(monkey.operation.arg)
				if err != nil {
					panic("not a number")
				}
			}

			if monkey.operation.oper == "*" {
				worryLevel *= arg
			} else if monkey.operation.oper == "+" {
				worryLevel += arg
			}

			// part2 : dividir 3
			worryLevel = worryLevel / 3

			// part3: Test
			if worryLevel%monkey.testDivisible == 0 {
				ms.monkeys[monkey.throwMonkeyIfTrue].items = append(ms.monkeys[monkey.throwMonkeyIfTrue].items, worryLevel)
			} else {
				ms.monkeys[monkey.throwMonkeyIfFalse].items = append(ms.monkeys[monkey.throwMonkeyIfFalse].items, worryLevel)
			}
		}
		ms.monkeys[monkeyID].items = []int{}
	}
}

func (ms *MonkeyState) monkeyBusiness() int {
	counts := []int{}
	for _, monkey := range ms.monkeys {
		counts = append(counts, monkey.itemsCount)
	}

	sort.Ints(counts)
	L := len(counts)
	return counts[L-1] * counts[L-2]
}

func (*Puzzle) Part1(input string) string {
	monkeyState := parseInput(input)
	for round := 1; round <= 20; round++ {
		monkeyState.processRound()
		// fmt.Printf("After round %d:\n", round)
		// for mID, monkey := range monkeyState.monkeys {
		// 	fmt.Printf("Monkey %d: %v\n", mID, monkey.items)
		// }
	}
	monkeyBusiness := monkeyState.monkeyBusiness()
	return fmt.Sprint(monkeyBusiness)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
