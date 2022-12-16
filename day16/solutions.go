package day16

import (
	"fmt"
	"sort"
	"strings"
)

type Puzzle struct{}

type Valve struct {
	label    string
	flowRate int
	isOpen   bool
	adjacent []string
}

type ValveState map[string]Valve

func parseInput(raw string) ValveState {
	valves := ValveState{}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		var label string
		var flowRate int
		lineParts := strings.Split(line, "; tunnels lead to valves ")
		if len(lineParts) == 1 {
			lineParts = strings.Split(line, "; tunnel leads to valve ")
		}

		fmt.Sscanf(lineParts[0], "Valve %s has flow rate=%d", &label, &flowRate)
		valves[label] = Valve{
			label:    label,
			flowRate: flowRate,
			adjacent: strings.Split(lineParts[1], ", "),
		}
	}
	return valves
}

func CloneStringSlice(ss []string) []string {
	newSlice := []string{}
	for _, s := range ss {
		newSlice = append(newSlice, strings.Clone(s))
	}
	return newSlice
}

func (v Valve) Clone() Valve {
	return Valve{
		label:    v.label,
		flowRate: v.flowRate,
		isOpen:   v.isOpen,
		adjacent: CloneStringSlice(v.adjacent),
	}
}

func (v ValveState) Clone() ValveState {
	vs := ValveState{}
	for label, valve := range v {
		vs[label] = valve.Clone()
	}
	return vs
}

func (v ValveState) CurrentState(curr string, minute int) string {
	openValves := []string{}
	for label, valve := range v {
		if valve.isOpen {
			openValves = append(openValves, label)
		}
	}
	sort.Slice(openValves, func(i, j int) bool { return openValves[i] < openValves[j] })
	return fmt.Sprintf("%d:%s:%s", minute, curr, strings.Join(openValves, ","))
}

type ValveCache struct {
	cache map[string]int
}

func (c ValveCache) findMaxPressureReleased(curr string, valvesState ValveState, minute int, pressureFlow, totalPressure int) int {
	currentState := valvesState.CurrentState(curr, minute)
	if x, exists := c.cache[currentState]; exists && x >= totalPressure {
		return c.cache[currentState]
	} else {
		c.cache[currentState] = totalPressure
	}

	// fmt.Println(curr, minute, totalPressure)
	if minute >= 30 {
		return totalPressure + pressureFlow
	}

	forkedResult := 0

	// fork opening Valve
	if !valvesState[curr].isOpen && valvesState[curr].flowRate > 0 {
		newState := valvesState.Clone()
		newValve := newState[curr]
		newValve.isOpen = true
		newState[curr] = newValve
		forkedResult = c.findMaxPressureReleased(curr, newState, minute+1, pressureFlow+newValve.flowRate, totalPressure+pressureFlow)
	}

	// fork for eachTunnel
	for _, nextValve := range valvesState[curr].adjacent {
		newState := valvesState.Clone()
		tmpForkedResult := c.findMaxPressureReleased(nextValve, newState, minute+1, pressureFlow, totalPressure+pressureFlow)
		if forkedResult < tmpForkedResult {
			forkedResult = tmpForkedResult
		}
	}
	return forkedResult
}

func (*Puzzle) Part1(input string) string {
	// parseInput()
	//   valve.flowRate
	//   valve.adjacent: []Valve
	//   valve.isOpen bool
	//   []Valve
	// simul()
	//   - entrar sala valvula X
	//   - "escolher" abrir ou não a valvula
	//   - "escolher" valvulas adjacentes
	//
	valves := parseInput(input)
	c := ValveCache{
		cache: make(map[string]int),
	}
	solution := c.findMaxPressureReleased("AA", valves, 1, 0, 0)
	return fmt.Sprint(solution)
}

//   eg: A-B, C-B
//     AA BB CC
// AA [   t   ]
// BB [ t    t]
// CC [   t   ]

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
