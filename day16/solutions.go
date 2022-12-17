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

func contains(ss []string, substr string) bool {
	for _, x := range ss {
		if x == substr {
			return true
		}
	}
	return false
}

func findNonEmptyValves(valves Valves) map[string]int {
	nonEmpty := map[string]int{}
	for label, flow := range valves.flowRates {
		if flow > 0 {
			nonEmpty[label] = flow
		}
	}
	return nonEmpty
}

type Queue []string

func (q *Queue) popLeft() string {
	elem := (*q)[0]
	*q = (*q)[1:]
	return elem
}

func computeDistanceToAll(connections map[string][]string, nonEmpty map[string]int, origin string) map[string]int {
	distances := map[string]int{}
	dist := 0
	q := append(Queue{}, connections[origin]...)
	for dist <= 26 {
		dist++
		tmpQueue := Queue{}
		for len(q) > 0 {
			next := q.popLeft()
			if _, ok := distances[next]; ok || next == origin {
				continue
			}
			if _, ok := nonEmpty[next]; ok {
				distances[next] = dist
			}
			for _, conn := range connections[next] {
				if !contains(tmpQueue, conn) {
					tmpQueue = append(tmpQueue, conn)
				}
			}
		}
		q = append(q, tmpQueue...)
	}
	return distances
}

func computeValvesDistances(connections map[string][]string, nonEmpty map[string]int) map[string]map[string]int {
	// {A : {B: 1, C: 2, D: 1, I: 1, ....}}
	distances := map[string]map[string]int{}
	for origin := range nonEmpty {
		distances[origin] = computeDistanceToAll(connections, nonEmpty, origin)
	}

	// must keep distances from first node
	distances["AA"] = computeDistanceToAll(connections, nonEmpty, "AA")

	return distances
}

type Cache struct {
	stateMaxPressure map[string]int
	distances        map[string]map[string]int
	valvesFlow       map[string]int
}

func pathToState(s []string) string {
	newS := append([]string{}, s[1:]...)
	sort.Strings(newS)
	return strings.Join(newS, ",")
}

func (c *Cache) calculateTotalFlow(path []string) (voidMinutes int, totalFlow int) {
	totalPressure := 0
	dist := 0

	// calculate total flow for this path
	for i, v := range path[1:] {
		neededMinutes := c.distances[path[i]][v] + 1
		dist += neededMinutes
		totalPressure += (30 - dist) * c.valvesFlow[v]
	}

	return 30 - dist, totalPressure
}

func (c *Cache) dfs(path []string, valveState string) int {
	currValve := path[len(path)-1]

	remaining, totalPressure := c.calculateTotalFlow(path)
	if n, ok := c.stateMaxPressure[valveState]; ok && n > totalPressure {
		return n
	}
	c.stateMaxPressure[valveState] = totalPressure

	maxPressure := totalPressure
	for nextValve := range c.distances[currValve] {
		if remaining < 0 { // cannot reach it in time
			break
		}
		if strings.Contains(valveState, nextValve) { // already open
			continue
		}

		newPath := append(path, nextValve)
		newState := pathToState(newPath)
		pressure := c.dfs(newPath, newState)
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}

	return maxPressure
}

func (*Puzzle) Part1(input string) string {
	valves := parseInput(input)
	nonEmptyValves := findNonEmptyValves(valves)
	distances := computeValvesDistances(valves.connections, nonEmptyValves)

	// possiblePaths := map[string]int{} // map[opened-valves]totalPressure
	cache := Cache{
		stateMaxPressure: make(map[string]int),
		distances:        distances,
		valvesFlow:       valves.flowRates,
	}
	solution := cache.dfs([]string{"AA"}, "")

	return fmt.Sprint(solution)
}
func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
