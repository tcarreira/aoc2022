package day19

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

const (
	Ore      int = 0
	Clay     int = 1
	Obsidian int = 2
	Geode    int = 3
)

type Blueprint [4][4]int

func parseInput(raw string) []Blueprint {
	bprints := []Blueprint{}
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		var bID, oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian int
		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &bID, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)

		bprint := Blueprint{
			{oreOre, 0, 0, 0},
			{clayOre, 0, 0, 0},
			{obsidianOre, obsidianClay, 0, 0},
			{geodeOre, 0, geodeObsidian, 0},
		}
		bprints = append(bprints, bprint)
	}
	return bprints
}

type State struct {
	maxMinutes int
	blueprint  Blueprint
	cache      map[string]int
	maxRate    [4]int
}

func max(a ...int) int {
	max := a[0]
	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

func setupState(b Blueprint) State {
	return State{
		maxMinutes: 24,
		blueprint:  b,
		cache:      make(map[string]int),
		maxRate: [4]int{
			max(b[Ore][Ore], b[Clay][Ore], b[Obsidian][Ore], b[Geode][Ore]),
			max(b[Ore][Clay], b[Clay][Clay], b[Obsidian][Clay], b[Geode][Clay]),
			max(b[Ore][Obsidian], b[Clay][Obsidian], b[Obsidian][Obsidian], b[Geode][Obsidian]),
			max(b[Ore][Geode], b[Clay][Geode], b[Obsidian][Geode], b[Geode][Geode]),
		},
	}
}

type Resources struct {
	ores     int
	clay     int
	obsidian int
	geodes   int
}

func (r Resources) Clone() Resources {
	return Resources{
		ores:     r.ores,
		clay:     r.clay,
		obsidian: r.obsidian,
		geodes:   r.geodes,
	}
}

func mineResources(resources, robots Resources) Resources {
	return Resources{
		ores:     resources.ores + robots.ores,
		clay:     resources.clay + robots.clay,
		obsidian: resources.obsidian + robots.obsidian,
		geodes:   resources.geodes + robots.geodes,
	}
}

func (st *State) dfs(minute int, resources, robots Resources) int {
	// fmt.Println(minute, resources, robots)
	if minute >= st.maxMinutes {
		return resources.geodes
	}

	// always make geode robot if possible
	if resources.ores >= st.blueprint[Geode][Ore] &&
		resources.obsidian >= st.blueprint[Geode][Obsidian] {
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.geodes++
		newResources.ores -= st.blueprint[Geode][Ore]
		newResources.obsidian -= st.blueprint[Geode][Obsidian]
		return st.dfs(minute+1, newResources, newRobots)
	}

	maxGeodes := 0

	if resources.ores >= st.blueprint[Obsidian][Ore] &&
		// make Obsidian robot
		resources.clay >= st.blueprint[Obsidian][Clay] &&
		robots.obsidian < st.maxRate[Obsidian] {
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.obsidian++
		newResources.ores -= st.blueprint[Obsidian][Ore]
		newResources.clay -= st.blueprint[Obsidian][Clay]
		newMaxGeodes := st.dfs(minute+1, newResources, newRobots)
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
	}

	if resources.ores >= st.blueprint[Ore][Ore] && robots.ores < st.maxRate[Ore] {
		// make ore robot
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.ores++
		newResources.ores -= st.blueprint[Ore][Ore]
		newMaxGeodes := st.dfs(minute+1, newResources, newRobots)
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
	}

	if resources.ores >= st.blueprint[Clay][Ore] && robots.clay < st.maxRate[Clay] {
		// make clay robot
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.clay++
		newResources.ores -= st.blueprint[Clay][Ore]
		newMaxGeodes := st.dfs(minute+1, newResources, newRobots)
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
	}

	resources = mineResources(resources, robots)
	newMaxGeodes := st.dfs(minute+1, resources, robots)
	if newMaxGeodes > maxGeodes {
		maxGeodes = newMaxGeodes
	}

	return maxGeodes
}

func (*Puzzle) Part1(input string) string {
	bprints := parseInput(input)

	solution := 0
	for id, bprint := range bprints {
		// fmt.Println("blueprint", id+1)
		st := setupState(bprint)
		geodes := st.dfs(0, Resources{}, Resources{ores: 1})

		qualityLevel := (id + 1) * geodes
		solution += qualityLevel
	}
	return fmt.Sprint(solution)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
