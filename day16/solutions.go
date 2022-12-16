package day16

import (
	"fmt"
	"sort"
	"strings"
)

type Puzzle struct{}

type Valves struct {
	flowRates   map[string]int
	connections map[string][]string
}

func parseInput(raw string) Valves {
	valves := Valves{
		flowRates:   make(map[string]int),
		connections: make(map[string][]string),
	}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		var label string
		var flowRate int
		var connections []string
		lineParts := strings.Split(line, "; tunnels lead to valves ")
		if len(lineParts) == 1 {
			lineParts = strings.Split(line, "; tunnel leads to valve ")
		}

		fmt.Sscanf(lineParts[0], "Valve %s has flow rate=%d", &label, &flowRate)
		connections = strings.Split(lineParts[1], ", ")

		valves.flowRates[label] = flowRate
		valves.connections[label] = connections

	}
	return valves
}

type ValvesState struct {
	current    string
	openValves []string
}

func (s ValvesState) CurrentState(minute int) string {
	sort.Strings(s.openValves)
	return fmt.Sprintf("%d:%s:%v", minute, s.current, s.openValves)
}
func (s ValvesState) Clone() ValvesState {
	openValves := []string{}
	openValves = append(openValves, s.openValves...)

	return ValvesState{
		current:    s.current,
		openValves: openValves,
	}
}

type ValveCache struct {
	cache  map[string]int
	valves Valves
}

func contains(s string, ss []string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

func (c ValveCache) findMaxPressureReleased(state ValvesState, minute, pressureFlow, totalPressure int) int {
	currentState := state.CurrentState(minute)
	if x, exists := c.cache[currentState]; exists && x >= totalPressure {
		return c.cache[currentState]
	} else {
		c.cache[currentState] = totalPressure
	}

	// fmt.Println(state.current, minute, totalPressure)
	if minute >= 30 {
		return totalPressure + pressureFlow
	}

	forkedResult := 0

	// fork opening Valve
	if !contains(state.current, state.openValves) && c.valves.flowRates[state.current] > 0 {
		newState := state.Clone()
		newState.openValves = append(newState.openValves, state.current)
		newPressureFlow := pressureFlow + c.valves.flowRates[state.current]
		forkedResult = c.findMaxPressureReleased(newState, minute+1, newPressureFlow, totalPressure+pressureFlow)
	}

	// fork for eachTunnel
	for _, nextValve := range c.valves.connections[state.current] {
		newState := state.Clone()
		newState.current = nextValve
		tmpForkedResult := c.findMaxPressureReleased(newState, minute+1, pressureFlow, totalPressure+pressureFlow)
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
		cache:  make(map[string]int),
		valves: valves,
	}

	state := ValvesState{current: "AA", openValves: []string{}}
	solution := c.findMaxPressureReleased(state, 1, 0, 0)
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
